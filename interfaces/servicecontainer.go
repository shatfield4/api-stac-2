/*
|--------------------------------------------------------------------------
| Service Container
|--------------------------------------------------------------------------
|
| This file performs the compiled dependency injection for your middlewares,
| controllers, services, providers, repositories, etc..
|
*/
package interfaces

import (
	"log"
	"os"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"api-grandtheftbacon/infrastructures/database/mysql"
	baconContract "api-grandtheftbacon/infrastructures/smartcontracts/bacon"
	theFryingPanContract "api-grandtheftbacon/infrastructures/smartcontracts/thefryingpan"
	nftRepository "api-grandtheftbacon/module/nft/infrastructure/repository"
	nftService "api-grandtheftbacon/module/nft/infrastructure/service"
	nftREST "api-grandtheftbacon/module/nft/interfaces/http/rest"
)

// ServiceContainerInterface contains the dependency injected instances
type ServiceContainerInterface interface {
	// REST
	RegisterNFTRESTCommandController() nftREST.NFTCommandController
	RegisterNFTRESTQueryController() nftREST.NFTQueryController
}

type kernel struct{}

var (
	m                            sync.Mutex
	k                            *kernel
	containerOnce                sync.Once
	mysqlDBHandler               *mysql.MySQLDBHandler
	EthHttpClient                *ethclient.Client
	EthWsClient                  *ethclient.Client
	BaconContractInstance        *baconContract.Bacon
	TheFryingPanContractInstance *theFryingPanContract.Thefryingpan
)

//================================= REST ===================================
// RegisterNFTRESTCommandController performs dependency injection to the RegisterNFTRESTCommandController
func (k *kernel) RegisterNFTRESTCommandController() nftREST.NFTCommandController {
	service := k.nftCommandServiceContainer()

	controller := nftREST.NFTCommandController{
		NFTCommandServiceInterface: service,
	}

	return controller
}

// RegisterNFTRESTQueryController performs dependency injection to the RegisterNFTRESTQueryController
func (k *kernel) RegisterNFTRESTQueryController() nftREST.NFTQueryController {
	service := k.nftQueryServiceContainer()

	controller := nftREST.NFTQueryController{
		NFTQueryServiceInterface: service,
	}

	return controller
}

//==========================================================================

func (k *kernel) nftCommandServiceContainer() *nftService.NFTCommandService {
	return NFTCommandServiceDI()
}

func (k *kernel) nftQueryServiceContainer() *nftService.NFTQueryService {
	repository := &nftRepository.NFTQueryRepository{
		MySQLDBHandlerInterface: mysqlDBHandler,
	}

	service := &nftService.NFTQueryService{
		NFTQueryRepositoryInterface: &nftRepository.NFTQueryRepositoryCircuitBreaker{
			NFTQueryRepositoryInterface: repository,
		},
		BaconContractInstance: BaconContractInstance,
	}

	return service
}

func NFTCommandServiceDI() *nftService.NFTCommandService {
	commandRepository := &nftRepository.NFTCommandRepository{
		MySQLDBHandlerInterface: mysqlDBHandler,
	}

	queryRepository := &nftRepository.NFTQueryRepository{
		MySQLDBHandlerInterface: mysqlDBHandler,
	}

	service := &nftService.NFTCommandService{
		NFTCommandRepositoryInterface: &nftRepository.NFTCommandRepositoryCircuitBreaker{
			NFTCommandRepositoryInterface: commandRepository,
		},
		NFTQueryRepositoryInterface: &nftRepository.NFTQueryRepositoryCircuitBreaker{
			NFTQueryRepositoryInterface: queryRepository,
		},
		BaconContractInstance: BaconContractInstance,
	}

	return service
}

func registerHandlers() {
	var err error

	/// connect to database
	mysqlDBHandler = &mysql.MySQLDBHandler{}
	err = mysqlDBHandler.Connect(os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"))
	if err != nil {
		log.Fatalf("[SERVER] mysql database is not responding %v", err)
	}

	/// connect to blockchain
	EthHttpClient, err = ethclient.Dial(os.Getenv("ETH_HTTP_ENDPOINT"))
	if err != nil {
		log.Fatal(err)
	}

	EthWsClient, err = ethclient.Dial(os.Getenv("ETH_WS_ENDPOINT"))
	if err != nil {
		log.Fatal(err)
	}

	/// load smart contracts
	BaconContractAddress = common.HexToAddress(os.Getenv("ETH_BACON_CONTRACT_ADDRESS"))
	BaconContractInstance, err = baconContract.NewBacon(BaconContractAddress, EthWsClient)
	if err != nil {
		log.Fatal(err)
	}

	TheFryingPanContractAddress = common.HexToAddress(os.Getenv("ETH_THEFRYINGPAN_CONTRACT_ADDRESS"))
	TheFryingPanContractInstance, err = theFryingPanContract.NewThefryingpan(TheFryingPanContractAddress, EthWsClient)
	if err != nil {
		log.Fatal(err)
	}

	/// load abi json
	BaconContractABI, err = abi.JSON(strings.NewReader(string(baconContract.BaconABI)))
	if err != nil {
		log.Fatal(err)
	}

	TheFryingPanContractABI, err = abi.JSON(strings.NewReader(string(theFryingPanContract.ThefryingpanABI)))
	if err != nil {
		log.Fatal(err)
	}

	// run event watcher
	go BaconContractEventWatcher()
	go TheFryingPanContractEventWatcher()
}

// ServiceContainer export instantiated service container once
func ServiceContainer() ServiceContainerInterface {
	m.Lock()
	defer m.Unlock()

	if k == nil {
		containerOnce.Do(func() {
			// register container handlers
			registerHandlers()

			k = &kernel{}
		})
	}
	return k
}
