package managers

import (
	"github.com/hjhcode/deploy-web/models"
)

func GetAllMirror() ([]map[string]interface{}, int) {
	mirrorList, err := models.Mirror{}.QueryAllMirror()
	if err != nil {
		panic(err.Error())
	}

	if mirrorList == nil {
		return nil, 0
	}

	var mirrorLists []map[string]interface{}
	for i := 0; i < len(mirrorList); i++ {
		mirrors := make(map[string]interface{})
		mirrors["id"] = mirrorList[i].Id
		mirrors["name"] = mirrorList[i].MirrorName + ":" + mirrorList[i].MirrorVersion
		mirrorLists = append(mirrorLists, mirrors)
	}
	count, err := models.Mirror{}.CountAllMirror()
	if err != nil {
		panic(err.Error())
	}

	return mirrorLists, int(count)
}

func getMirrorName(mirrorId int64) string {
	mirror, err := models.Mirror{}.GetById(mirrorId)
	if err != nil {
		panic(err.Error())
	}
	if mirror == nil {
		return ""
	}

	return mirror.MirrorName + ":" + mirror.MirrorVersion
}

func getMirrorDescribe(mirrorId int64) string {
	mirror, err := models.Mirror{}.GetById(mirrorId)
	if err != nil {
		panic(err.Error())
	}
	if mirror == nil {
		return ""
	}
	return mirror.MirrorDescribe
}

func getMirrorNameById(mirrorId int64) string {
	if mirrorId == 0 {
		return ""
	}
	mirror, err := models.Mirror{}.GetById(mirrorId)
	if err != nil {
		panic(err.Error())
	}
	if mirror == nil {
		return ""
	}

	return mirror.MirrorName
}
