package managers

import (
	"github.com/hjhcode/deploy-web/models"
)

func getMirrorName(mirrorId int64) string {
	mirror, err := models.Mirror{}.GetById(mirrorId)
	if err != nil {
		panic(err.Error())
	}
	if mirror == nil {
		return ""
	}

	return mirror.MirrorName
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
