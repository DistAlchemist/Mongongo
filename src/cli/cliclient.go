package cli

type CliClient struct {
	mgClient mongongo.Client
}

func (cc *CliClient) executeQueryOnServer(line string) {
	//
	res := cc.mgClient.executeQuery(line)
}
