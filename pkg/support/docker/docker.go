package docker

import (
    "fmt"
    "regexp"
)

var dockerImageRegex = regexp.MustCompile(`([^/]*)/([^:]*):(.*)`)

type DockerImageInfo struct {
    ImageName   string
    ImageTag    string
    ImageOrg    string
    ImageFqname string
}

func ParseDockerImageUrl(imageUrl string) (dii DockerImageInfo, err error) {
    match := dockerImageRegex.FindStringSubmatch(imageUrl)

    if len(match) == 0 {
        return DockerImageInfo{}, fmt.Errorf("error pasing the docker url")
    } else {
        dii.ImageOrg = match[1]
        dii.ImageName = match[2]
        dii.ImageTag = match[3]
        dii.ImageFqname = imageUrl
    }

    return dii, nil
}