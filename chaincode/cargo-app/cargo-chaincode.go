/*
 * The smart contract for Cargo and Containers:
 * Track & Trace of Cargo and Containers
 * Container States --> Available, Loaded, In-Cargo, In-Transit, Unloaded, Customs Pending.
 * Cargo States --> Ready, In-Transit, Arrived.
 * Participant Roles --> Container Supplier, Transporter, Exporter, Importer, Customs Officer
 */

package main

/* Imports
 * 2 utility libraries for formatting, and reading and writing JSON
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the Cargo, Container structure, with 10 properties.  Structure tags are used by encoding/json library
type Cargo struct {
	HashId string `json:"hashId"`
	TxnId string `json:"txnId"`
	Timestamp string `json:"timestamp"`
	CargoId string `json:"cargoId"`
	ShippedFrom string `json:"shippedFrom"`
	ShippedTo string `json:"shippedTo"`
	CargoLocation string `json:"cargoLocation"`
	TransportationType string `json:"transportationType"`
	ContainerQty string `json:"containerQty"`
	Owner string `json:"owner"`
	AssociatedContainerHashIds[] string `json:"associatedContainerHashIds"`
	Status string `json:"status"`
}

type Container struct {
	HashId string `json:"hashId"`
	Timestamp string `json:"timestamp"`
	Manufacturer string `json:"manufacturer"`
	Status string `json:"status"`
	LoadedItems string `json:"loadedItems"`
	Owner string `json:"owner"`
	CargoId string `json:"cargoId"`
	CustomClearanceStatus string `json:"customClearanceStatus"`
	ShippedFrom string `json:"shippedFrom"`
	ShippedTo string `json:"shippedTo"`
	ContainerLocation string `json:"containerLocation"`
}

type Participant struct {
	HashId string `json:"hashId"`
	Name string `json:"name"`
	EmailId string `json:"emailId"`
	Role string `json:"role"`
}


/*
 * The Init method is called when the Smart Contract "fabcargo" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "fabcargo"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "registerParticipant" {  		        // Done - This is to add legitimate users with role in system.
		return s.registerParticipant(APIstub, args)
	} else if function == "addNewContainer" {     			// Done - This is for granting a new container to Transporter by Container Supplier.
		return s.addNewContainer(APIstub, args)
	} else if function == "getParticipant" {
		return s.getParticipant(APIstub, args)
	} else if function == "loadContainerWithPackages" {   	// Done - This is to support container loading process(only available containers can be loaded).
		return s.loadContainerWithPackages(APIstub, args)
	} else if function == "createCargoLoadContainers" {		// Done - This is to create a new Cargo and associate loaded containers with it.
		return s.createCargoLoadContainers(APIstub, args)
	} else if function == "updateCargoAttributes" { 	    // Done - This is to support cargo movement and IOT sensor based updates. 
		return s.updateCargoAttributes(APIstub, args)
	} else if function == "updateCargoCoordinates" {		// Done - This is to update cargo and its associated container coordinates (IOT based).
		return s.updateCargoCoordinates(APIstub, args)
	} else if function == "changeCargoCustody" {		    // Done - This is to record cargo ownership changes.
		return s.changeCargoCustody(APIstub, args)
	} else if function == "updateContainerAttributes" {     // Done - This is to support container IOT sensor based updates.
		return s.updateContainerAttributes(APIstub, args)
	} else if function == "changeContainerCustody" {    	// Done - This is to record container ownership changes.
		return s.changeContainerCustody(APIstub, args)
	} else if function == "unloadContainerFromCargo" {      // Done - This is to support container unloading process.
		return s.unloadContainerFromCargo(APIstub, args)
	} else if function == "traceCargo" {                    // Done -  This is to trace Cargo
		return s.traceCargo(APIstub, args)
	} else if function == "trackCargoDetails" {			    // Done - This is to track Cargo
		return s.trackCargoDetails(APIstub, args)
	} else if function == "traceContainer" {                // Done - This is to trace Container
		return s.traceContainer(APIstub, args)
	} else if function == "trackContainerDetails" {			// Done - This is to track Container
		return s.trackContainerDetails(APIstub, args)
	} else if function == "getLoadedContainers" {			// Done - This is to get all Loaded state Containers
		return s.getLoadedContainers(APIstub)
	} else if function == "getAvilableContainers" {			// Done - This is to get all Available state Containers
		return s.getAvilableContainers(APIstub)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) registerParticipant(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	participantAsBytes, errr := APIstub.GetState(args[0])

	if errr != nil {
		return shim.Error("Failed to get participant: "+errr.Error())	
	} else if participantAsBytes != nil {
		return shim.Error("This User already exists. Name: "+args[1]+", EmailId: "+args[2])		
	}

	var participant = Participant{HashId: args[0], Name: args[1], EmailId: args[2], Role: args[3]}

	participantNewAsBytes, _ := json.Marshal(participant)
	APIstub.PutState(args[0], participantNewAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) addNewContainer(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 11 {
		return shim.Error("Incorrect number of arguments. Expecting 11")
	}

	//containerAsBytes, errr := APIstub.GetState(args[0])
	
	//ownerAsBytes, err := APIstub.GetState(args[6])
	
	//participant := Participant{}
	
	//json.Unmarshal(ownerAsBytes, &participant)
	

	//if errr != nil {
	//	return shim.Error("Failed to get container: "+errr.Error())	
	//} else if containerAsBytes != nil {
	//	return shim.Error("This Container record already exists. HashId: "+args[0])		
	//} else if err != nil {
	//	return shim.Error("Failed to get owner: "+err.Error())	
	//} else if participant.Role != 'Transporter' {
	//	return shim.Error("New Owner is not a Transporter")
	//}
	
	if args[6] != "Transporter" {
		return shim.Error("New Owner is not a Transporter")
	}

	var container = Container{HashId: args[0], Timestamp: args[1], Manufacturer: args[2], Status: args[3], LoadedItems: args[4], Owner: args[5],CargoId: args[6], CustomClearanceStatus: args[7], ShippedFrom: args[8], ShippedTo: args[9], ContainerLocation: args[10]}

	containerAsBytes, _ := json.Marshal(container)
	APIstub.PutState(args[0], containerAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) getParticipant(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	containerAsBytes, _ := APIstub.GetState(args[0])
	
	return shim.Success(containerAsBytes)
}

func (s *SmartContract) loadContainerWithPackages(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 8")
	}

	containerAsBytes, _ := APIstub.GetState(args[0])
	container := Container{}

	json.Unmarshal(containerAsBytes, &container)
	if container.Status != "Available" {
		return shim.Error("Container is not Available for loading packages.")
	}
	container.Timestamp = args[1]
	container.Status = args[2]
	container.LoadedItems = args[3]
	container.CustomClearanceStatus = args[4]
	container.ShippedFrom = args[5]
	container.ShippedTo = args[6]
	container.ContainerLocation = args[7]

	containerAsBytes, _ = json.Marshal(container)
	APIstub.PutState(args[0], containerAsBytes)
	
	return shim.Success(nil)
}

func (s *SmartContract) createCargoLoadContainers(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 12 {
		return shim.Error("Incorrect number of arguments. Expecting 12")
	}
	
	// Check for Cargo In-Transit status
	////cargoCheckAsBytes, _ := APIstub.GetState(args[0])
	////cargoCheck := Cargo{}

	////json.Unmarshal(cargoCheckAsBytes, &cargoCheck)
	////if cargoCheck.Status == 'In-Transit' && cargoCheck.CargoLocation != args[6] {  
		// Call recordContainerAttributes to update Location
		
		////cargoAsBytes, _ = json.Marshal(cargo)
		////APIstub.PutState(args[0], cargoAsBytes)
	////}

	var ids []string = strings.Split(args[10],",")
	
	var cargo = Cargo{HashId: args[0], TxnId: args[1], Timestamp: args[2], CargoId: args[3], ShippedFrom: args[4], ShippedTo: args[5], CargoLocation: args[6], TransportationType: args[7], ContainerQty: args[8], Owner: args[9], AssociatedContainerHashIds: ids, Status: args[11]}

	cargoAsBytes, _ := json.Marshal(cargo)
	APIstub.PutState(args[0], cargoAsBytes)
	
	for index, containerHashId := range ids {
		
		containerAsBytes, _ := APIstub.GetState(containerHashId)
		container := Container{}

		json.Unmarshal(containerAsBytes, &container)
		container.CargoId = args[0]
		container.Timestamp = args[2]
		container.Status = "InCargo"
		container.ContainerLocation = args[6]
		
		containerAsBytes, _ = json.Marshal(container)
		
		APIstub.PutState(containerHashId, containerAsBytes)
	}
	
	return shim.Success(nil)
}


func (s *SmartContract) updateCargoAttributes(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	
	if len(args) != 9 {
		return shim.Error("Incorrect number of arguments. Expecting 9")
	}

	cargoAsBytes, err := APIstub.GetState(args[0])
	cargo := Cargo{}

	if err != nil {
		return shim.Error("Failed to get Cargo: "+err.Error())
	}

	var ids []string = strings.Split(args[7],",")

	json.Unmarshal(cargoAsBytes, &cargo)
	cargo.TxnId = args[1]
	cargo.Timestamp = args[2]
	cargo.ShippedFrom = args[3]
	cargo.ShippedTo = args[4]
	cargo.TransportationType = args[5]
	cargo.ContainerQty = args[6]
	cargo.AssociatedContainerHashIds = ids
	cargo.Status = args[8]

	cargoAsBytes, _ = json.Marshal(cargo)
	APIstub.PutState(args[0], cargoAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) updateCargoCoordinates(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	cargoAsBytes, err := APIstub.GetState(args[0])
	cargo := Cargo{}

	if err != nil {
		return shim.Error("Failed to get Cargo: "+err.Error())
	}

	json.Unmarshal(cargoAsBytes, &cargo)
	cargo.Timestamp = args[1]
	cargo.CargoLocation = args[2]

	cargoAsBytes, _ = json.Marshal(cargo)
	APIstub.PutState(args[0], cargoAsBytes)


	for index, containerHashId := range cargo.AssociatedContainerHashIds {
		
		containerAsBytes, _ := APIstub.GetState(containerHashId)
		container := Container{}

		json.Unmarshal(containerAsBytes, &container)
		container.ContainerLocation = args[1]
		container.ContainerLocation = args[2]
		
		containerAsBytes, _ = json.Marshal(container)
		APIstub.PutState(containerHashId, containerAsBytes)
	}
	return shim.Success(nil)
}

func (s *SmartContract) changeCargoCustody(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	cargoAsBytes, _ := APIstub.GetState(args[0])
	cargo := Cargo{}

	json.Unmarshal(cargoAsBytes, &cargo)
	cargo.Owner = args[1]

	cargoAsBytes, _ = json.Marshal(cargo)
	APIstub.PutState(args[0], cargoAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) updateContainerAttributes(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 9 {
		return shim.Error("Incorrect number of arguments. Expecting 9")
	}

	containerAsBytes, err := APIstub.GetState(args[0])
	container := Container{}

	if err != nil {
		return shim.Error("Failed to get Container: "+err.Error())
	}

	json.Unmarshal(containerAsBytes, &container)
	container.Timestamp = args[1]
	container.Manufacturer = args[2]
	container.Status = args[3]
	container.LoadedItems = args[4]
	container.CustomClearanceStatus = args[5]
	container.ShippedFrom = args[6]
	container.ShippedTo = args[7]
	container.ContainerLocation = args[8]

	containerAsBytes, _ = json.Marshal(container)
	APIstub.PutState(args[0], containerAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) changeContainerCustody(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	containerAsBytes, err := APIstub.GetState(args[0])
	container := Container{}

	if err != nil {
		return shim.Error("Failed to get Container: "+err.Error())
	}

	json.Unmarshal(containerAsBytes, &container)
	container.Owner = args[1]

	containerAsBytes, _ = json.Marshal(container)
	APIstub.PutState(args[0], containerAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) unloadContainerFromCargo(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	cargoAsBytes, _ := APIstub.GetState(args[0])
	cargo := Cargo{}

	json.Unmarshal(cargoAsBytes, &cargo)
	
	for index, containerHashId := range cargo.AssociatedContainerHashIds {
	
		if containerHashId == args[1] {
			// Remove the element at index index from cargo.AssociatedContainerHashIds.
			cargo.AssociatedContainerHashIds[index] = cargo.AssociatedContainerHashIds[len(cargo.AssociatedContainerHashIds)-1] // Copy last element to index i.
			cargo.AssociatedContainerHashIds[len(cargo.AssociatedContainerHashIds)-1] = ""   // Erase last element (write zero value).
			cargo.AssociatedContainerHashIds = cargo.AssociatedContainerHashIds[:len(cargo.AssociatedContainerHashIds)-1]   // Truncate slice.
			break;
		}
	}
	
	cargoAsBytes, _ = json.Marshal(cargo)
	APIstub.PutState(args[0], cargoAsBytes)
	
	// Update Container Status as 'Unloaded'
	
	containerAsBytes, _ := APIstub.GetState(args[1])
	container := Container{}

	json.Unmarshal(containerAsBytes, &container)
	container.Status = "Unloaded"
	
	containerAsBytes, _ = json.Marshal(container)
	APIstub.PutState(args[1], containerAsBytes)
	
	return shim.Success(nil)
}

func (s *SmartContract) traceCargo(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	type TraceCargo struct {
		TxId    string   `json:"txId"`
		Value   Cargo   `json:"value"`
	}
	var traceCargo []TraceCargo;
	var cargo Cargo

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	cargoId := args[0]
	fmt.Printf("- start getTraceForCargo: %s\n", cargoId)

	// Get Trace
	resultsIterator, err := APIstub.GetHistoryForKey(cargoId)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		traceCargoData, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		var tx TraceCargo
		tx.TxId = traceCargoData.TxId                     //copy transaction id over
		json.Unmarshal(traceCargoData.Value, &cargo)     //un stringify it aka JSON.parse()
		if traceCargoData.Value == nil {                  //cargo has been deleted
			var emptyCargo Cargo
			tx.Value = emptyCargo                 //copy nil cargo
		} else {
			json.Unmarshal(traceCargoData.Value, &cargo) //un stringify it aka JSON.parse()
			tx.Value = cargo                      //copy cargo over
		}
		traceCargo = append(traceCargo, tx)              //add this tx to the list
	}
	fmt.Printf("- getTraceForCargo returning:\n%s", traceCargo)

	//change to array of bytes
	traceCargoAsBytes, _ := json.Marshal(traceCargo)     //convert to array of bytes
	return shim.Success(traceCargoAsBytes)
}

func (s *SmartContract) trackCargoDetails(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	cargoAsBytes, err := APIstub.GetState(args[0])
	if err != nil {
		return shim.Error("Cargo doesn't exist. "+err.Error())
	}

	return shim.Success(cargoAsBytes)
}

func (s *SmartContract) traceContainer(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	type TraceContainer struct {
		TxId    string   `json:"txId"`
		Value   Container   `json:"value"`
	}
	var traceContainer []TraceContainer;
	var container Container

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	containerId := args[0]
	fmt.Printf("- start getTraceForContainer: %s\n", containerId)

	// Get Trace
	resultsIterator, err := APIstub.GetHistoryForKey(containerId)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		traceContainerData, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		var tx TraceContainer
		tx.TxId = traceContainerData.TxId                     //copy transaction id over
		json.Unmarshal(traceContainerData.Value, &container)     //un stringify it aka JSON.parse()
		if traceContainerData.Value == nil {                  //cargo has been deleted
			var emptyContainer Container
			tx.Value = emptyContainer                 //copy nil cargo
		} else {
			json.Unmarshal(traceContainerData.Value, &container) //un stringify it aka JSON.parse()
			tx.Value = container                      //copy ontainer over
		}
		traceContainer = append(traceContainer, tx)              //add this tx to the list
	}
	fmt.Printf("- getTraceForContainer returning:\n%s", traceContainer)

	//change to array of bytes
	traceContainerAsBytes, _ := json.Marshal(traceContainer)     //convert to array of bytes
	return shim.Success(traceContainerAsBytes)
}

func (s *SmartContract) trackContainerDetails(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	containerAsBytes, err := APIstub.GetState(args[0])
	if err != nil {
		return shim.Error("Container doesn't exist. "+err.Error())
	}

	return shim.Success(containerAsBytes)
}

func (s *SmartContract) getLoadedContainers(APIstub shim.ChaincodeStubInterface) sc.Response {
	
	type LoadedContainers struct {
		Containers   []Container   `json:"containers"`
	}
	var loadedContainers LoadedContainers
	
	//containersIterator, err := APIstub.GetStateByRange("o0", "o9999999999999999999")
	containersIterator, err := APIstub.GetStateByRange("1001", "9999999")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer containersIterator.Close()

	for containersIterator.HasNext() {
		aKeyValue, err := containersIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		queryKeyAsStr := aKeyValue.Key
		queryValAsBytes := aKeyValue.Value
		fmt.Println("on owner id - ", queryKeyAsStr)
		var container Container
		json.Unmarshal(queryValAsBytes, &container)                   //un stringify it aka JSON.parse()

		if container.Status == "Loaded" {                                        //only return Loaded container
			loadedContainers.Containers = append(loadedContainers.Containers, container)  //add this container to the list
		}
	}
	fmt.Println("container array - ", loadedContainers.Containers)

	//change to array of bytes
	loadedContainersAsBytes, _ := json.Marshal(loadedContainers)              //convert to array of bytes
	return shim.Success(loadedContainersAsBytes)
}

func (s *SmartContract) getAvilableContainers(APIstub shim.ChaincodeStubInterface) sc.Response {
	type AvilableContainers struct {
		Containers   []Container   `json:"containers"`
	}
	var avilableContainers AvilableContainers
	
	//containersIterator, err := APIstub.GetStateByRange("o0", "o9999999999999999999")
	containersIterator, err := APIstub.GetStateByRange("1001", "9999999")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer containersIterator.Close()

	for containersIterator.HasNext() {
		aKeyValue, err := containersIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		queryKeyAsStr := aKeyValue.Key
		queryValAsBytes := aKeyValue.Value
		fmt.Println("on owner id - ", queryKeyAsStr)
		var container Container
		json.Unmarshal(queryValAsBytes, &container)                   //un stringify it aka JSON.parse()

		if container.Status == "Available" {                                        //only return Loaded container
			avilableContainers.Containers = append(avilableContainers.Containers, container)  //add this container to the list
		}
	}
	fmt.Println("container array - ", avilableContainers.Containers)

	//change to array of bytes
	avilableContainersAsBytes, _ := json.Marshal(avilableContainers)              //convert to array of bytes
	return shim.Success(avilableContainersAsBytes)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
