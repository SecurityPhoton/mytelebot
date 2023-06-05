pipeline {
    agent any
    environment {
        REPO = 'https://github.com/pontarr/mytelebot'
        BRANCH = 'develop'
    }
    parameters {

        choice(name: 'OS', choices: ['linux', 'darwin', 'windows', 'all'], description: 'Pick OS')
        
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

                    script {
                                env.OS = params.OS
                                env.ARCH = params.ARCH
                            }
                    echo 'MAKE BUILD'
                    echo 'envOS = ${env.OS} envARCH = ${env.ARCH}'
                    sh 'make build'
                  }
        }
    }
}
