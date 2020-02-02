APP_NAME = 'home-cms'
COMP_NAME = "backend"
IMAGE_TAG = ''
IMAGE_URI = ''
IMAGE_LATEST = ''
WAR = ''
FILENAME = ''


DOCKER_REGISTRY_HOST_NAME = "192.168.11.3:10000"
DOCKER_REGISTRY_ORG = "ci"
DOCKER_REGISTRY_CREDIENTIAL_ID = "harbor"

GIT_REGISTRY = "https://github.com/michaelanony/home-cms.git"
COMPILE_AGENT_IMAGE = ""


/** BRANCH = "" is Defined in Parameter  */

COMPILE_SOURCE = true


pipeline {
    agent {
    label 'jenkins-ci'
  }

    options { timestamps () }
    stages {



        stage('Git Checkout') {
            steps {
                dir('source'){
                    git branch: "${BRANCH}", credentialsId: '1afda173-d7ab-4ffb-89cc-36106a82febe', url: "${GIT_REGISTRY}"

                    script {
                        COMMIT_HASH = sh(returnStdout: true, script: 'git rev-parse HEAD').trim().take(7)
                        BRANCH_MOD = BRANCH.replaceAll(/\//,'_')
                        def now = new Date()
                        DATE = now.format("yyMMdd", TimeZone.getTimeZone('UTC'))
                        IMAGE_TAG="${APP_NAME}:${COMP_NAME}.${BRANCH_MOD}.${COMMIT_HASH}.${DATE}.${BUILD_NUMBER}"
                        IMAGE_URI="${DOCKER_REGISTRY_HOST_NAME}/${DOCKER_REGISTRY_ORG}/${IMAGE_TAG}"
                        IMAGE_LATEST="${DOCKER_REGISTRY_HOST_NAME}/${DOCKER_REGISTRY_ORG}/${APP_NAME}"
                        echo "${IMAGE_URI}"
                    }
                }
            }
        }

         stage('Compile') {
            when {
                expression { COMPILE_SOURCE == true }
            }
            steps {
                script {
                    sh "cd source && go build -o server"
                }
            }
        }

        stage('Docker Image Build') {
            steps{
                script{
                    sh 'cd source'
                    sh 'ls -l'
                    docker.withRegistry("http://${DOCKER_REGISTRY_HOST_NAME}", "${DOCKER_REGISTRY_CREDIENTIAL_ID}") {

                        def image = docker.build("${IMAGE_URI}","-f source/Dockerfile --force-rm .")
                        image.push()
                        image.push('latest.${COMP_NAME}')
                        }
                    }
                }
            }

        stage('Deploy dev') {
                when {
                    // Only deploy if it is requested
                    expression { params.DEPLOY == 'DEV' && params.DEPLOY_OR_NOT == true}
                }
                steps{
                script{
                    withKubeConfig([credentialsId: 'ec258f6e-c4ae-463e-b132-58a2541232a6', serverUrl: 'https://kubernetes.default.svc.cluster.local']) {
                        if (env.APPLICATION != 'Hub Mongoconnector') {
                            sh "kubectl set image -n ${APP_NAME}-${ENVIRONMENT} --record deployment deploy-${APP_NAME}-${COMP_NAME} ${APP_NAME}-${COMP_NAME}=${IMAGE_URI}"
                            sh "kubectl rollout status deploy deploy-${APP_NAME}-${COMP_NAME} -n ${APP_NAME}-${ENVIRONMENT}"
                        }else{
                            sh "kubectl set image -n db-elastic deployment deploy-mongo-connector-mc1 mongo-connector=${IMAGE_URI}"
                            sh "#echo 'hello'"
                        }
                    }
                }
            }
        }




        stage ("Clean") {
            steps {
                dir('source') {
                    deleteDir()
                }

                dir('source_tmp') {
                    deleteDir()
                }
            }
        }
    }

}