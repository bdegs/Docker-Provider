//go:build windows

package extension

import (
	winio "github.com/Microsoft/go-winio"
)

func getExtensionConfigResponse(jsonBytes []byte) (string, error) {
	pipePath := "\\\\.\\\\pipe\\\\CAgentStream_CloudAgentInfo_AzureMonitorAgent"
	config_namedpipe, err := winio.DialPipe(pipePath, nil)
	if err != nil {
		logger.Printf("AMA::extensionconfig::getExtensionConfigResponse:: error opening pipe %s: %v", pipePath, err)
		return "", err
	}
	defer config_namedpipe.Close()
	number_bytes, err := config_namedpipe.Write(jsonBytes)
	if err != nil {
		logger.Printf("AMA::extensionconfig::getExtensionConfigResponse:: write error: %v", err)
		return "", err
	}

	read_buffer := make([]byte, ReadBufferSize)
	number_bytes, err = config_namedpipe.Read(read_buffer)
	if err != nil {
		logger.Printf("AMA::extensionconfig::getExtensionConfigResponse:: read error: %v", err)
		return "", err
	}
	read_buffer = read_buffer[:number_bytes]
	response := string(read_buffer)

	return response, nil
}
