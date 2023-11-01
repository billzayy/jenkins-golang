pipeline {
  agent any

  tools { go '1.21.3' }
    
  stages {
    stage('Git Checkout') {
        steps {
            checkout scmGit(branches: [[name: '*/main']], extensions: [], userRemoteConfigs: [[url: 'https://github.com/billzayy/jenkins-golang']])
        }
    }
    stage('Test Version') {
      steps {
        sh 'go version'
      }
    }
    stage('Run'){
        steps{
            sh 'go build . && ./jenkins-golang'
        }
    }
  }
}