pipeline {
    agent any
    tools {
        go 'go1.14'
    }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
        EXECPATH = "${env.WORKSPACE}" + "/" + "nagv"
    }
    stages {        
        stage('Pre Test') {
            steps {
                echo 'Installing dependencies'
                sh 'go version'
                sh 'go get -u golang.org/x/lint/golint'
            }
        }
        
        stage('Build') {
            steps {
                echo 'Compiling and building'
                sh 'go build'
            }
        }

        stage('Test') {
            steps {
                withEnv(["PATH+GO=${GOPATH}/bin"]){
                    echo 'Running vetting'
                    sh 'go vet .'
                    echo 'Running linting'
                    sh 'golint .'
                    //echo 'Running test'
                    //sh 'cd test && go test -v'
                    //sh 'ls -l $EXECPATH'
                }
            }
        }
        stage('Deploy') {
            steps {
               withCredentials([sshUserPrivateKey(credentialsId: "jvssh", keyFileVariable: 'keyfile')]) {
          sh  """
          scp -oStrictHostKeyChecking=no -i ${keyfile} $EXECPATH jenkinsuser@v141008m-adm:/var/tmp 
          """
          }    
         }
        }
 
    }
    post {
        always {
            emailext body: "${currentBuild.currentResult}: Job ${env.JOB_NAME} build ${env.BUILD_NUMBER}\n More info at: ${env.BUILD_URL}",
                recipientProviders: [[$class: 'DevelopersRecipientProvider'], [$class: 'RequesterRecipientProvider']],
                to: "${params.RECIPIENTS}",
                subject: "Jenkins Build ${currentBuild.currentResult}: Job ${env.JOB_NAME}"
            
        }
    }  
}
