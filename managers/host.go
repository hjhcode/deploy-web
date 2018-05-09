package managers

import "github.com/hjhcode/deploy-web/models"

func AddNewHost(hostName string, hostIp string) (bool, string) {
	result := checkHostName(hostName)
	if !result {
		return false, "机器名已存在"
	}

	host := &models.Host{
		HostName: hostName,
		Ip:       hostIp,
	}

	_, err := models.Host{}.Add(host)
	if err != nil {
		panic(err.Error())
	}

	return true, ""
}

func DelHost(hostId int64) bool {
	err := models.Host{}.Remove(hostId)
	if err != nil {
		panic(err.Error())
	}

	return true
}

func UpdateHost(hostId int64, hostName string, hostIp string) bool {
	host := &models.Host{
		Id:       hostId,
		HostName: hostName,
		Ip:       hostIp,
	}

	err := models.Host{}.Update(host)
	if err != nil {
		panic(err.Error())
	}

	return true
}

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

func checkHostName(hostName string) bool {
	host, err := models.Host{}.GetByHostName(hostName)
	if err != nil {
		panic(err.Error())
	}
	if host != nil {
		return false
	}
	return true
}
