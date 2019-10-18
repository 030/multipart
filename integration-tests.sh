#!/bin/bash -eux

TOOL="${1:-./go-multipart}"

validate(){
    if [ -z "$TOOL" ]; then
        echo "No deliverable defined. Assuming that 'go run main.go' 
should be run."
        TOOL="go run main.go"
    fi
}

nexus(){
    docker run --rm -d -p 9999:8081 --name nexus sonatype/nexus3:3.16.1
}

readiness(){
    until docker logs nexus | grep 'Started Sonatype Nexus OSS'
    do
        echo "Nexus unavailable"
        sleep 10
    done
}

artifacts(){
    $TOOL -url \
    http://localhost:9999/service/rest/v1/components?repository=maven-releases \
    -user admin -pass admin123 -F \
    "maven2.asset1=@utils/test-files-multipart/file1.pom,\
    maven2.asset1.extension=pom,\
    maven2.asset2=@utils/test-files-multipart/file1.jar,\
    maven2.asset2.extension=jar,\
    maven2.asset3=@utils/test-files-multipart/file1-sources.jar,\
    maven2.asset3.extension=sources.jar"
}

cleanup(){
    docker stop nexus
}

main(){
    validate
    nexus
    readiness
    artifacts
    cleanup
}

main