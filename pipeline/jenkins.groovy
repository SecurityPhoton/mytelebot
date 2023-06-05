pipeline {
    agent any
    environment {
        REPO = 'https://github.com/pontarr/mytelebot'
        BRANCH = 'develop'
        PATH = "/usr/local/go/bin:${env.PATH}"
    }
    parameters {

        choice(name: 'OS', choices: ['linux', 'darwin', 'windows', 'darwin', 'all'], description: 'Pick OS')
        
        choice(name: 'ARCH', choices: ['amd64', 'arm64', '386', 'arm'], description: 'Pick ARCH')

    }
    stages {
        stage("clone") {
            steps {
                echo 'Clone Repository'
                git branch: "${BRANCH}", url: "${REPO}"
                
                  }
         }
        
        stage("test") {
            steps {
                echo 'MAKE TEST'
                sh 'make test'
            }
        }

        stage("build") {
            steps {
                    echo 'MAKE BUILD:'
                    sh "make build TARGETOS=${params.OS} TARGETARCH=${params.ARCH}"
                  }
        }

        stage("image") {
            steps {
                    echo 'MAKE IAMGE:'
                    sh "make image"
                }
            }

        stage("push"){
            steps {
                echo 'MAKE PUSH'
                script {
                    docker.withRegistry('','ghcr.io') {
                        sh 'make push'
                    }
                }

            }
        }
    }
}
