package main

import "testing"


func TestDecodeTransaction(t *testing.T){
	var transactionString="020000000001010ccc140e766b5dbc884ea2d780c5e91e4eb77597ae64288a42575228b79e234900000000000000000002bd37060000000000225120245091249f4f29d30820e5f36e1e5d477dc3386144220bd6f35839e94de4b9cae81c00000000000016001416d31d7632aa17b3b316b813c0a3177f5b6150200140838a1f0f1ee607b54abf0a3f55792f6f8d09c3eb7a9fa46cd4976f2137ca2e3f4a901e314e1b827c3332d7e1865ffe1d7ff5f5d7576a9000f354487a09de44cd00000000";
	expectedJson :=`{"Version":2,"Inputs":[{"PreviousTxHash":"0ccc140e766b5dbc884ea2d780c5e91e4eb77597ae64288a42575228b79e2349","PreviousTxOutIndex":0,"ScriptSig":"","Sequence":0}],"Outputs":[{"Value":407485,"ScriptPubKey":"5120245091249f4f29d30820e5f36e1e5d477dc3386144220bd6f35839e94de4b9ca"},{"Value":7400,"ScriptPubKey":"001416d31d7632aa17b3b316b813c0a3177f5b615020"}],"Locktime":0}`
	result,err:= decodeTransaction(transactionString);
	if err != nil {
         t.Errorf("Error found")
        
    }

	if string(result)!=expectedJson{
		 t.Errorf("Expected JSON does not match actual JSON")
	}

}