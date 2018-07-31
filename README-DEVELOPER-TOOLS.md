# Quick Start - Creation from JSON file 

For some situations it may be easier to have the configuration represented as JSON rather than conifguring individually. In this scenario you can either build the JSON file yourself or monitor the API POST call for the JSON data sent to CCP. This can be achieved using the browsers built in developer tools. 

If using the developer tools process you will first need to create the resource required in the GUI. This will allow you to find the correct JSON structure and reuse this structure in any future calls.

### 1. Login to CCP and navigate to the resource you are trying to create. In this example we will create a new cluster

![alt tag](https://github.com/conmurphy/ccp-clientlibrary-go/blob/master/images/cluster_creation.png)

### 2. Open the developer tools window within your browser and navigate to the network tab

### 3. Fill in the required details

### 4. Select `Save` and you should see a success banner indicating the request was submitted and the nodes are now being provisioning

![alt tag](https://github.com/conmurphy/ccp-clientlibrary-go/blob/master/images/successfull_creation.png)

### 5. Look through the developer tools window to find the call made to CCP. In this example it is `clusters/`. 

### 6. You should see in the right hand widow that the `Request Method` is `POST`

### 7. Scroll to the bottom and look at the `Request Payload`. You can select `View source` to obtain the data in JSON format

![alt tag](https://github.com/conmurphy/ccp-clientlibrary-go/blob/master/images/developer_tools.png)

### 8. Save the JSON into a file 

![alt tag](https://github.com/conmurphy/ccp-clientlibrary-go/blob/master/images/json.png)

### 9. You can now parse this JSON file with the library
