package main
	
import (
	"context"
	"crypto/tls"
	"errors"
	"net"
	"net/http"
	"runtime/debug"
	"strconv"

	// "/auth"
	"flag"
	"fmt"
	"os"

	"code.it4i.cz/lexis/wp8/dataset-management-interface.git/server/datasetManager"
	"code.it4i.cz/lexis/wp8/dataset-management-interface.git/server/statusManager"
	//	"code.it4i.cz/lexis/wp8/dataset-management-interface.git/server/dbManager"
	//	"code.it4i.cz/lexis/wp8/dataset-management-interface.git/server/irodsManager"
	"code.it4i.cz/lexis/wp8/dataset-management-interface.git/restapi"
	"github.com/spf13/viper"
	l "gitlab.com/cyclops-utilities/logging"
	"github.com/Nerzal/gocloak/v7"
)

type BasicUserData struct {
	UserSub string
}

type generalConfig struct {
	LogFile      string
	LogToConsole bool
	LogLevel     string
	ServerPort   int
}

type irodsConfig struct {
	Service string
	Gridmap string
	Sshfs string
	Steering string
	Compress string
	Datasize string
	Encryption string
}

type keycloakConfig struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Realm        string `json:"realm"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	UseHttp      bool   `json:"use_http"`
	RedirectURL  string `json:"redirect_url"`
}

type configuration struct {
	General  generalConfig
	Irods    irodsConfig
	Keycloak keycloakConfig `json:"keycloak"`
}

var (
	cfg     configuration
	version string
)

func parseConfig() (c configuration) {

	c = configuration{
		General: generalConfig{
			LogFile:      viper.GetString("general.logfile"),
			LogToConsole: viper.GetBool("general.logtoconsole"),
			LogLevel:     viper.GetString("general.loglevel"),
			ServerPort:   viper.GetInt("general.serverport"),
		},
		Irods: irodsConfig{
			Service:	viper.GetString("irods.service"),
			Gridmap:	viper.GetString("irods.gridmap"),
			Sshfs:		viper.GetString("irods.sshfs"),
			Steering:	viper.GetString("irods.steering"),
			Compress:	viper.GetString("irods.compress"),
			Datasize:	viper.GetString("irods.datasize"),
                        Encryption:	viper.GetString("irods.encryption"),
		},
		Keycloak: keycloakConfig{
			Host:         viper.GetString("keycloak.host"),
			Port:         viper.GetInt("keycloak.port"),
			Realm:        viper.GetString("keycloak.realm"),
			ClientID:     viper.GetString("keycloak.clientid"),
			ClientSecret: viper.GetString("keycloak.clientsecret"),
			UseHttp:      viper.GetBool("keycloak.usehttp"),
			RedirectURL:  viper.GetString("keycloak.redirecturl"),
		},
	}
	return
}

// getKeycloaktService returns the keycloak service; note that there has to be exceptional
// handling of port 80 and port 443
func getKeycloakService(c keycloakConfig) (s string) {
	if c.UseHttp {
		s = "http://" + c.Host
		if c.Port != 80 {
			s = s + ":" + strconv.Itoa(c.Port)
		}
	} else {
		s = "https://" + c.Host
		if c.Port != 443 {
			s = s + ":" + strconv.Itoa(c.Port)
		}
	}
	return
}

// this is quite useless rght now..
func Authorizer(r *http.Request) (e error) {
	l.Warning.Printf("In Authorizer: Warning - this is not yet implemented...\n")
	debug.PrintStack()
	return nil
}

// AuthKeycloak is called whenever it is necessary to check if a token which is
// presented to access an API is valid
func AuthKeycloak(token string, scopes []string) (i interface{}, returnErr error) {

	//l.Debug.Printf("AuthKeycloak: Performing authentication check - token = %v...\n", token)

	keycloakService := getKeycloakService(cfg.Keycloak)
	client := gocloak.NewClient(keycloakService)

	ctx := context.Background()

	client.RestyClient().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	_, err := client.LoginClient(ctx, cfg.Keycloak.ClientID, cfg.Keycloak.ClientSecret, cfg.Keycloak.Realm)
	if err != nil {
		l.Warning.Printf("Error logging in to keycloak - %v\n", err.Error())
		returnErr = errors.New("Unable to log in to keycloak")
		return
	}

	u, err := client.GetUserInfo(ctx, token, cfg.Keycloak.Realm)
	if err != nil {
		l.Warning.Printf("Error getting user Info: %v\n", err.Error())
		returnErr = errors.New("Unable to get UserInfo from keycloak: "+err.Error())
		return
	} else {
		if u != nil {
			i = BasicUserData{
				UserSub: *u.Sub,
			}
			return
		}
	}
	return
}

// init function - reads in configuration file and creates logger
func init() {

	confFile := flag.String("conf", "./config", "configuration file path (without toml extension)")
	flag.Parse()

	//placeholder code as the default value will ensure this situation will never arise
	if len(*confFile) == 0 {
		fmt.Println("Usage: dataset-management-interface --conf=/path/to/configuration/file")
		os.Exit(0)
	}

	// err := gcfg.ReadFileInto(&cfg, *confFile)
	viper.SetConfigName(*confFile) // name of config file (without extension)
	viper.SetConfigType("toml")
	viper.AddConfigPath(".") // path to look for the config file in
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Fatal error reading config file - err %s\nExiting...\n", err)
		os.Exit(1)
	}

	cfg = parseConfig()

	l.InitLogger(cfg.General.LogFile, cfg.General.LogLevel, cfg.General.LogToConsole)

	l.Info.Printf("Dataset Management Interface version %v initialized", version)
}

// main function creates the database connection and launches the endpoint handlers
func main() {
	// Initiate business logic implementers.
	// This is the main function, so here the implementers' dependencies can be
	// injected, such as database, parameters from environment variables, or different
	// clients for different APIs.
	mon := statusManager.New()

	d := datasetManager.New(cfg.Irods.Service, cfg.Irods.Gridmap, 
		cfg.Irods.Sshfs, cfg.Irods.Steering, cfg.Irods.Compress, cfg.Irods.Datasize, cfg.Irods.Encryption, mon,
		cfg.Keycloak.RedirectURL)
	// Initiate the http handler, with the objects that are implementing the business logic.
	h, err := restapi.Handler(restapi.Config{
		Logger:       l.Info.Printf,
		AuthKeycloak: AuthKeycloak,
		DataSetManagementAPI: d,
		StagingAPI: d,
		StatusManagementAPI:       mon,
	})
	if err != nil {
		l.Error.Printf("Error creating REST handler: %s ...Exiting\n", err)
		os.Exit(2)
	}

	// note that this runs on all interfaces right now
	serviceLocation := ":" + strconv.Itoa(cfg.General.ServerPort)

	l.Info.Printf("Starting to serve data-management-service, access server on http://localhost%v\n", serviceLocation)

	// Run the standard http server on ipv4
	listener, err := net.Listen("tcp4", serviceLocation)
	if err != nil {
		l.Error.Printf ("%s", err)
		os.Exit(3)
	}
	l.Error.Printf((http.Serve(listener, h)).Error())
}
