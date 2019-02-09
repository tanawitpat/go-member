pipeline {
    agent {
        label "master"
    }
    options {
        skipDefaultCheckout true
    }
    stages {
        stage("PULL CODE") {
            steps {
                git branch: "DevMaster", credentialsId: "3cffc65557880f217286fc6e244382d7", url: "https://github.com/OuiPTN/go-member.git"
                script{
                    gitCommit = sh (script: "git rev-parse HEAD | cut -c1-7", returnStdout: true).trim()
                    currentBuild.displayName = "#${env.BUILD_NUMBER}_${gitCommit}"
                }
            }
        }
        stage("BUILD IMAGE") {
            steps {
                sh "docker build . -t go-member"
            }
        }
        stage("DEPLOY") {
            steps {
                sh """
                docker run --rm -d --network go-member --name go-member \
                -p 45000:8050 \
                go-member
                """
            }
        }
    }
}