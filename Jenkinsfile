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
    stage('Check Docker'){
        steps{
            sh 'docker version'
        }
    }
    stage('Build & Push Images'){
        steps{
            script{
                withDockerRegistry(credentialsId: 'billzay-docker', toolName: 'Docker', url: 'https://index.docker.io/v1/') {
                    sh 'docker build -t coderbillzay/jenkins-golang .'
                    sh 'docker push coderbillzay/jenkins-golang'
                }
            }
        }
    }
    stage('Run Docker-Compose'){
        steps{
            sh 'docker-compose up -d'
        }
    }
  }
}