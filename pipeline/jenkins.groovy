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
        
        stage("build") {
            steps {
                    echo 'MAKE BUILD'
                    sh 'make build OS=${params.OS} ARCH=${params.ARCH}'
                  }
        }
    }
}
