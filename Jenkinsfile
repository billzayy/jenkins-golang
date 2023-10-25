pipeline {
  agent any

  tools { go '1.21.3' }
    
    environment {
        GIT_BRANCH = 'main'
        GIT_URL = 'git@github.com:billzayy/jenkins-golang.git'
    }
    
  stages {
    stage('Git Checkout') {
        steps {
            git branch: "${GIT_BRANCH}", url: "${GIT_URL}", poll: false
        }
    }
    stage('Test Version') {
      steps {
        sh 'go version'
      }
    }
    stage('Run'){
        steps{
            sh 'go run .'
        }
    }
  }
}