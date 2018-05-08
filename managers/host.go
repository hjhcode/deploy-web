package managers

import "github.com/hjhcode/deploy-web/models"

func GetAllHost() ([]*models.Host, int) {
	hostList, err := models.Host{}.QueryAllHost()
	if err != nil {
		panic(err.Error())
	}

	if hostList == nil {
		return nil, 0
	}

	count, err := models.Host{}.CountAllHost()
	if err != nil {
		panic(err.Error())
	}

	return hostList, int(count)
}

func getHostName(hostId int64) string {
	host, err := models.Host{}.GetById(hostId)
	if err != nil {
		panic(err.Error())
	}
	if host == nil {
		return ""
	}

	return host.HostName
}
