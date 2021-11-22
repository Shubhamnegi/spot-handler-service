def slackUser = "@shubham negi"
def slackChannel = "#jenkins"
// def testCommand = "go test"
def testCommand = "ls"
def jenkins
node {
   fileLoader.withGit('https://github.com/LimeTray/jenkins-pipeline-scripts.git', 'master', 'limetray-github', '') {
       jenkins = fileLoader.load('jenkins');
   }
}
jenkins.start(testCommand, slackUser, slackChannel)