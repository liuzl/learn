# step by step

1. `mvn archetype:generate -DgroupId=maven.demo.start -DartifactId=HelloMaven -DarchetypeArtifactId=maven-archetype-quickstart -DinteractiveMode=false`

2. `cd HelloMaven && mvn package`

3. `java -cp target/HelloMaven-1.0-SNAPSHOT.jar maven.demo.start.App`
