make buildpipeline {
    agent any
    environment {
        REPO = 'https://github.com/pontarr/mytelebot'
        BRACH = 'develop'
    }
    parameters {

        choice(name: 'OS', choices: ['linux', 'darwin', 'windows', 'all'], description: 'Pick OS')
        
        choice(name: 'ARCH', choices: ['amd64', 'arm64', '386', 'arm'], description: 'Pick ARCH')

    }
    stages {
        stage('clone') {
            steps {
                echo "Clone Repository ${params.OS}"
                git branch: "${BRANCH}", url: "${REPO}"
                
                echo "Build for arch: ${params.ARCH}"

                 }
         }
        
        stage('build') {
            echo "MAKE BUILD"
            script {
                    env.OS = params.OS
                    env.ARCH = params.ARCH
                }
            
            sh 'make build OS=${env.OS} ARCH=${env.ARCH}'
        }
    }
}
