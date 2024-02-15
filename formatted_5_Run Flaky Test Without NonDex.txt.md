2024-02-14T18:27:10.4387100Z ##[group]Run ./cicd/run-without-nondex
2024-02-14T18:27:10.4388158Z [36;1m./cicd/run-without-nondex[0m
2024-02-14T18:27:10.5375042Z shell: /usr/bin/bash -e {0}
2024-02-14T18:27:10.5376302Z env:
2024-02-14T18:27:10.5378161Z   MAVEN_OPTS: -Dorg.slf4j.simpleLogger.log.org.apache.maven.plugins.shade=error
2024-02-14T18:27:10.5381629Z   GOROOT: /home/fedja/Desktop/my_runner/actions-runner/_work/_tool/go/1.21.7/x64
2024-02-14T18:27:10.5385426Z   JAVA_HOME: /home/fedja/Desktop/my_runner/actions-runner/_work/_tool/Java_Temurin-Hotspot_jdk/11.0.22-7/x64
2024-02-14T18:27:10.5388277Z   TODAY: 20240214
2024-02-14T18:27:10.5389374Z   YESTERDAY: 20240213
2024-02-14T18:27:10.5390476Z ##[endgroup]
2024-02-14T18:27:10.5691992Z 2024/02/14 19:27:10 Running command: mvn -B test -f pom.xml -e -pl v1 -Dcheckstyle.skip -Dspotless.check.skip -Dtest=com.google.cloud.teleport.templates.common.DatastoreConvertersTest#testCheckNoKeyAllInvalid
2024-02-14T18:27:12.1780538Z [**INFO**] Error stacktraces are turned on.
2024-02-14T18:27:12.2221209Z [**INFO**] Scanning for projects...
2024-02-14T18:27:13.6027903Z [**INFO**] ------------------------------------------------------------------------
2024-02-14T18:27:13.6030314Z [**INFO**] Detecting the operating system and CPU architecture
2024-02-14T18:27:13.6037514Z [**INFO**] ------------------------------------------------------------------------
2024-02-14T18:27:13.6044138Z [**INFO**] os.detected.name: linux
2024-02-14T18:27:13.6046143Z [**INFO**] os.detected.arch: x86_64
2024-02-14T18:27:13.6049114Z [**INFO**] os.detected.bitness: 64
2024-02-14T18:27:13.6052111Z [**INFO**] os.detected.version: 6.5
2024-02-14T18:27:13.6054864Z [**INFO**] os.detected.version.major: 6
2024-02-14T18:27:13.6057772Z [**INFO**] os.detected.version.minor: 5
2024-02-14T18:27:13.6080115Z [**INFO**] os.detected.release: ubuntu
2024-02-14T18:27:13.6082017Z [**INFO**] os.detected.release.version: 22.04
2024-02-14T18:27:13.6084873Z [**INFO**] os.detected.release.like.ubuntu: true
2024-02-14T18:27:13.6087377Z [**INFO**] os.detected.release.like.debian: true
2024-02-14T18:27:13.6090335Z [**INFO**] os.detected.classifier: linux-x86_64
2024-02-14T18:27:13.6187001Z [**INFO**] 
2024-02-14T18:27:13.6191192Z [**INFO**] ------------< com.google.cloud.teleport:classic-templates >-------------
2024-02-14T18:27:13.6194235Z [**INFO**] Building Google Cloud Teleport Classic 1.0-SNAPSHOT
2024-02-14T18:27:13.6196176Z [**INFO**] --------------------------------[ jar ]---------------------------------
2024-02-14T18:27:14.1048279Z [**INFO**] Downloading from splunk-artifactory: https://splunk.jfrog.io/splunk/ext-releases-local/com/google/cloud/teleport/metadata/metadata/1.0-SNAPSHOT/maven-metadata.xml
2024-02-14T18:27:14.1054444Z [**INFO**] Downloading from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/metadata/metadata/1.0-SNAPSHOT/maven-metadata.xml
2024-02-14T18:27:15.0273574Z [**INFO**] Downloaded from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/metadata/metadata/1.0-SNAPSHOT/maven-metadata.xml (789 B at 851 B/s)
2024-02-14T18:27:15.0445921Z [**INFO**] Downloading from splunk-artifactory: https://splunk.jfrog.io/splunk/ext-releases-local/com/google/cloud/teleport/metadata/metadata/1.0-SNAPSHOT/metadata-1.0-20240109.172859-10.pom
2024-02-14T18:27:15.5568559Z [**INFO**] Downloading from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/metadata/metadata/1.0-SNAPSHOT/metadata-1.0-20240109.172859-10.pom
2024-02-14T18:27:15.8278809Z [**INFO**] Downloaded from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/metadata/metadata/1.0-SNAPSHOT/metadata-1.0-20240109.172859-10.pom (3.3 kB at 12 kB/s)
2024-02-14T18:27:15.8354352Z [**INFO**] Downloading from splunk-artifactory: https://splunk.jfrog.io/splunk/ext-releases-local/com/google/cloud/teleport/templates/1.0-SNAPSHOT/maven-metadata.xml
2024-02-14T18:27:15.8360802Z [**INFO**] Downloading from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/templates/1.0-SNAPSHOT/maven-metadata.xml
2024-02-14T18:27:16.1006666Z [**INFO**] Downloaded from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/templates/1.0-SNAPSHOT/maven-metadata.xml (609 B at 2.3 kB/s)
2024-02-14T18:27:16.3522306Z [**INFO**] Downloading from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/templates/1.0-SNAPSHOT/templates-1.0-20240109.172859-10.pom
2024-02-14T18:27:16.9004813Z [**INFO**] Downloaded from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/templates/1.0-SNAPSHOT/templates-1.0-20240109.172859-10.pom (21 kB at 39 kB/s)
2024-02-14T18:27:19.4458499Z [**INFO**] Downloading from splunk-artifactory: https://splunk.jfrog.io/splunk/ext-releases-local/com/google/cloud/teleport/it-google-cloud-platform/1.0-SNAPSHOT/maven-metadata.xml
2024-02-14T18:27:19.4461385Z [**INFO**] Downloading from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/it-google-cloud-platform/1.0-SNAPSHOT/maven-metadata.xml
2024-02-14T18:27:19.7240074Z [**INFO**] Downloaded from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/it-google-cloud-platform/1.0-SNAPSHOT/maven-metadata.xml (793 B at 2.9 kB/s)
2024-02-14T18:27:19.8338944Z [**INFO**] Downloading from splunk-artifactory: https://splunk.jfrog.io/splunk/ext-releases-local/com/google/cloud/teleport/it-google-cloud-platform/1.0-SNAPSHOT/it-google-cloud-platform-1.0-20240109.172859-5.pom
2024-02-14T18:27:20.2081877Z [**INFO**] Downloading from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/it-google-cloud-platform/1.0-SNAPSHOT/it-google-cloud-platform-1.0-20240109.172859-5.pom
2024-02-14T18:27:20.4697461Z [**INFO**] Downloaded from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/it-google-cloud-platform/1.0-SNAPSHOT/it-google-cloud-platform-1.0-20240109.172859-5.pom (6.4 kB at 25 kB/s)
2024-02-14T18:27:20.4763730Z [**INFO**] Downloading from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/integration-testing-lib/1.0-SNAPSHOT/maven-metadata.xml
2024-02-14T18:27:20.4773859Z [**INFO**] Downloading from splunk-artifactory: https://splunk.jfrog.io/splunk/ext-releases-local/com/google/cloud/teleport/integration-testing-lib/1.0-SNAPSHOT/maven-metadata.xml
2024-02-14T18:27:20.7421559Z [**INFO**] Downloaded from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/integration-testing-lib/1.0-SNAPSHOT/maven-metadata.xml (621 B at 2.3 kB/s)
2024-02-14T18:27:20.8621493Z [**INFO**] Downloading from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/integration-testing-lib/1.0-SNAPSHOT/integration-testing-lib-1.0-20240109.172859-7.pom
2024-02-14T18:27:21.1464357Z [**INFO**] Downloaded from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/integration-testing-lib/1.0-SNAPSHOT/integration-testing-lib-1.0-20240109.172859-7.pom (6.2 kB at 22 kB/s)
2024-02-14T18:27:21.5369752Z [**INFO**] Downloading from splunk-artifactory: https://splunk.jfrog.io/splunk/ext-releases-local/com/google/cloud/teleport/it-datadog/1.0-SNAPSHOT/maven-metadata.xml
2024-02-14T18:27:21.5377428Z [**INFO**] Downloading from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/it-datadog/1.0-SNAPSHOT/maven-metadata.xml
2024-02-14T18:27:21.8038134Z [**INFO**] Downloaded from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/it-datadog/1.0-SNAPSHOT/maven-metadata.xml (779 B at 2.9 kB/s)
2024-02-14T18:27:22.0418740Z [**INFO**] Downloading from splunk-artifactory: https://splunk.jfrog.io/splunk/ext-releases-local/com/google/cloud/teleport/it-datadog/1.0-SNAPSHOT/it-datadog-1.0-20240109.172859-4.pom
2024-02-14T18:27:22.4240582Z [**INFO**] Downloading from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/it-datadog/1.0-SNAPSHOT/it-datadog-1.0-20240109.172859-4.pom
2024-02-14T18:27:22.6822338Z [**INFO**] Downloaded from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/it-datadog/1.0-SNAPSHOT/it-datadog-1.0-20240109.172859-4.pom (2.6 kB at 9.9 kB/s)
2024-02-14T18:27:23.5641102Z [**INFO**] Downloading from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/metadata/metadata/1.0-SNAPSHOT/metadata-1.0-20240109.172859-10.jar
2024-02-14T18:27:23.5645495Z [**INFO**] Downloading from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/it-google-cloud-platform/1.0-SNAPSHOT/it-google-cloud-platform-1.0-20240109.172859-5.jar
2024-02-14T18:27:23.5649490Z [**INFO**] Downloading from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/it-datadog/1.0-SNAPSHOT/it-datadog-1.0-20240109.172859-4.jar
2024-02-14T18:27:23.9492112Z [**INFO**] Downloaded from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/metadata/metadata/1.0-SNAPSHOT/metadata-1.0-20240109.172859-10.jar (41 kB at 105 kB/s)
2024-02-14T18:27:24.1919212Z [**INFO**] Downloaded from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/it-datadog/1.0-SNAPSHOT/it-datadog-1.0-20240109.172859-4.jar (23 kB at 37 kB/s)
2024-02-14T18:27:24.5378850Z [**INFO**] Downloaded from ossrh: https://oss.sonatype.org/content/repositories/snapshots/com/google/cloud/teleport/it-google-cloud-platform/1.0-SNAPSHOT/it-google-cloud-platform-1.0-20240109.172859-5.jar (245 kB at 251 kB/s)
2024-02-14T18:27:24.5641333Z [**INFO**] 
2024-02-14T18:27:24.5650381Z [**INFO**] --- maven-enforcer-plugin:3.0.0-M1:enforce (enforce) @ classic-templates ---
2024-02-14T18:27:29.7377332Z [**WARNING**] Failure to find io.confluent:kafka-avro-serializer:jar:5.3.2 in https://splunk.jfrog.io/splunk/ext-releases-local was cached in the local repository, resolution will not be reattempted until the update interval of splunk-artifactory has elapsed or updates are forced
2024-02-14T18:27:29.7380531Z 
2024-02-14T18:27:29.7381101Z Try downloading the file manually from the project website.
2024-02-14T18:27:29.7381953Z 
2024-02-14T18:27:29.7382230Z Then, install it using the command: 
2024-02-14T18:27:29.7384389Z     mvn install:install-file -DgroupId=io.confluent -DartifactId=kafka-avro-serializer -Dversion=5.3.2 -Dpackaging=jar -Dfile=/path/to/file
2024-02-14T18:27:29.7385640Z 
2024-02-14T18:27:29.7386190Z Alternatively, if you host your own repository you can deploy the file there: 
2024-02-14T18:27:29.7388772Z     mvn deploy:deploy-file -DgroupId=io.confluent -DartifactId=kafka-avro-serializer -Dversion=5.3.2 -Dpackaging=jar -Dfile=/path/to/file -Durl=[url] -DrepositoryId=[id]
2024-02-14T18:27:29.7390348Z 
2024-02-14T18:27:29.7390566Z Path to dependency: 
2024-02-14T18:27:29.7391483Z 	1) com.google.cloud.teleport:classic-templates:jar:1.0-SNAPSHOT
2024-02-14T18:27:29.7392776Z 	2) org.apache.beam:beam-runners-google-cloud-dataflow-java:jar:2.53.0
2024-02-14T18:27:29.7394012Z 	3) org.apache.beam:beam-sdks-java-io-kafka:jar:2.53.0
2024-02-14T18:27:29.7395039Z 	4) io.confluent:kafka-avro-serializer:jar:5.3.2
2024-02-14T18:27:29.7395606Z 
2024-02-14T18:27:29.7395618Z 
2024-02-14T18:27:29.7396023Z   io.confluent:kafka-avro-serializer:jar:5.3.2
2024-02-14T18:27:29.7396544Z 
2024-02-14T18:27:29.7397553Z from the specified remote repositories:
2024-02-14T18:27:29.7399067Z   splunk-artifactory (https://splunk.jfrog.io/splunk/ext-releases-local, releases=true, snapshots=true),
2024-02-14T18:27:29.7400917Z   ossrh (https://oss.sonatype.org/content/repositories/snapshots, releases=true, snapshots=true),
2024-02-14T18:27:29.7402540Z   central (https://repo.maven.apache.org/maven2, releases=true, snapshots=false)
2024-02-14T18:27:29.7403393Z 
2024-02-14T18:27:29.7486091Z [**WARNING**] Failure to find io.confluent:kafka-schema-registry-client:jar:5.3.2 in https://splunk.jfrog.io/splunk/ext-releases-local was cached in the local repository, resolution will not be reattempted until the update interval of splunk-artifactory has elapsed or updates are forced
2024-02-14T18:27:29.7488856Z 
2024-02-14T18:27:29.7490072Z Try downloading the file manually from the project website.
2024-02-14T18:27:29.7490670Z 
2024-02-14T18:27:29.7490908Z Then, install it using the command: 
2024-02-14T18:27:29.7493079Z     mvn install:install-file -DgroupId=io.confluent -DartifactId=kafka-schema-registry-client -Dversion=5.3.2 -Dpackaging=jar -Dfile=/path/to/file
2024-02-14T18:27:29.7494513Z 
2024-02-14T18:27:29.7495055Z Alternatively, if you host your own repository you can deploy the file there: 
2024-02-14T18:27:29.7497780Z     mvn deploy:deploy-file -DgroupId=io.confluent -DartifactId=kafka-schema-registry-client -Dversion=5.3.2 -Dpackaging=jar -Dfile=/path/to/file -Durl=[url] -DrepositoryId=[id]
2024-02-14T18:27:29.7499510Z 
2024-02-14T18:27:29.7499722Z Path to dependency: 
2024-02-14T18:27:29.7500567Z 	1) com.google.cloud.teleport:classic-templates:jar:1.0-SNAPSHOT
2024-02-14T18:27:29.7501798Z 	2) org.apache.beam:beam-runners-google-cloud-dataflow-java:jar:2.53.0
2024-02-14T18:27:29.7503003Z 	3) org.apache.beam:beam-sdks-java-io-kafka:jar:2.53.0
2024-02-14T18:27:29.7503998Z 	4) io.confluent:kafka-schema-registry-client:jar:5.3.2
2024-02-14T18:27:29.7504575Z 
2024-02-14T18:27:29.7504590Z 
2024-02-14T18:27:29.7505066Z   io.confluent:kafka-schema-registry-client:jar:5.3.2
2024-02-14T18:27:29.7505653Z 
2024-02-14T18:27:29.7505902Z from the specified remote repositories:
2024-02-14T18:27:29.7507384Z   splunk-artifactory (https://splunk.jfrog.io/splunk/ext-releases-local, releases=true, snapshots=true),
2024-02-14T18:27:29.7509278Z   ossrh (https://oss.sonatype.org/content/repositories/snapshots, releases=true, snapshots=true),
2024-02-14T18:27:29.7511029Z   central (https://repo.maven.apache.org/maven2, releases=true, snapshots=false)
2024-02-14T18:27:29.7511978Z 
2024-02-14T18:27:32.8162490Z [**INFO**] 
2024-02-14T18:27:32.8164371Z [**INFO**] --- maven-enforcer-plugin:3.0.0-M1:enforce (enforce-banned-dependencies) @ classic-templates ---
2024-02-14T18:27:33.0512580Z [**INFO**] 
2024-02-14T18:27:33.0514628Z [**INFO**] --- jacoco-maven-plugin:0.8.8:prepare-agent (default-prepare-agent) @ classic-templates ---
2024-02-14T18:27:33.1576490Z [**INFO**] argLine set to -javaagent:/home/fedja/.m2/repository/org/jacoco/org.jacoco.agent/0.8.8/org.jacoco.agent-0.8.8-runtime.jar=destfile=/home/fedja/Desktop/my_runner/actions-runner/_work/PrivateCITest/PrivateCITest/v1/target/jacoco.exec,excludes=/*InformationSchemaScanner.*:/*AutoValue_*
2024-02-14T18:27:33.1580038Z [**INFO**] 
2024-02-14T18:27:33.1581176Z [**INFO**] --- protobuf-maven-plugin:0.6.1:compile (default) @ classic-templates ---
2024-02-14T18:27:33.5695562Z [**INFO**] Compiling 2 proto file(s) to /home/fedja/Desktop/my_runner/actions-runner/_work/PrivateCITest/PrivateCITest/v1/target/generated-sources/protobuf/java
2024-02-14T18:27:33.5882542Z [**INFO**] 
2024-02-14T18:27:33.5883958Z [**INFO**] --- protobuf-maven-plugin:0.6.1:compile-custom (default) @ classic-templates ---
2024-02-14T18:27:33.7727254Z [**INFO**] Compiling 2 proto file(s) to /home/fedja/Desktop/my_runner/actions-runner/_work/PrivateCITest/PrivateCITest/v1/target/generated-sources/protobuf/grpc-java
2024-02-14T18:27:33.7826764Z [**INFO**] 
2024-02-14T18:27:33.7828195Z [**INFO**] --- avro-maven-plugin:1.11.3:schema (default) @ classic-templates ---
2024-02-14T18:27:34.0139354Z [**WARNING**] Ignored the com.google.cloud.teleport.bigtable.BigtableCell.timestamp.logicalType property ("timestamp-micros"). It should probably be nested inside the "type" for the field.
2024-02-14T18:27:34.2771050Z [**INFO**] 
2024-02-14T18:27:34.2772350Z [**INFO**] --- maven-resources-plugin:2.6:resources (default-resources) @ classic-templates ---
2024-02-14T18:27:34.3411566Z [**INFO**] Using 'UTF-8' encoding to copy filtered resources.
2024-02-14T18:27:34.3444341Z [**INFO**] Copying 5 resources
2024-02-14T18:27:34.3474595Z [**INFO**] Copying 2 resources
2024-02-14T18:27:34.3476309Z [**INFO**] Copying 2 resources
2024-02-14T18:27:34.3479556Z [**INFO**] 
2024-02-14T18:27:34.3481279Z [**INFO**] --- maven-compiler-plugin:3.11.0:compile (default-compile) @ classic-templates ---
2024-02-14T18:27:34.4378584Z [**INFO**] Changes detected - recompiling the module! :dependency
2024-02-14T18:27:34.4646324Z [**INFO**] Compiling 164 source files with javac [debug target 11] to target/classes
2024-02-14T18:27:42.6416887Z [**WARNING**] /home/fedja/Desktop/my_runner/actions-runner/_work/PrivateCITest/PrivateCITest/v1/src/main/java/com/google/cloud/teleport/splunk/HttpEventPublisher.java:[108,19] [AutoValueMutable] An @AutoValue property that is a primitive array returns the original array, which can therefore be modified by the caller. If this is OK, you can suppress this warning with @SuppressWarnings("mutable"). Otherwise, you should replace the property with an immutable type, perhaps a simple wrapper around the original array.
2024-02-14T18:27:42.6426656Z [**WARNING**] Implicitly compiled files were not subject to annotation processing.
2024-02-14T18:27:42.6430158Z   Use -proc:none to disable annotation processing or -implicit to specify a policy for implicit compilation.
2024-02-14T18:27:42.6436596Z [**INFO**] /home/fedja/Desktop/my_runner/actions-runner/_work/PrivateCITest/PrivateCITest/v1/target/generated-sources/protobuf/java/com/google/cloud/teleport/spanner/proto/ExportProtos.java: Some input files use or override a deprecated API.
2024-02-14T18:27:42.6444885Z [**INFO**] /home/fedja/Desktop/my_runner/actions-runner/_work/PrivateCITest/PrivateCITest/v1/target/generated-sources/protobuf/java/com/google/cloud/teleport/spanner/proto/ExportProtos.java: Recompile with -Xlint:deprecation for details.
2024-02-14T18:27:42.6453778Z [**INFO**] /home/fedja/Desktop/my_runner/actions-runner/_work/PrivateCITest/PrivateCITest/v1/src/main/java/com/google/cloud/teleport/templates/DatastoreToText.java: Some input files use unchecked or unsafe operations.
2024-02-14T18:27:42.6461894Z [**INFO**] /home/fedja/Desktop/my_runner/actions-runner/_work/PrivateCITest/PrivateCITest/v1/src/main/java/com/google/cloud/teleport/templates/DatastoreToText.java: Recompile with -Xlint:unchecked for details.
2024-02-14T18:27:42.6465454Z [**INFO**] 
2024-02-14T18:27:42.6467153Z [**INFO**] --- maven-resources-plugin:2.6:testResources (default-testResources) @ classic-templates ---
2024-02-14T18:27:42.6469402Z [**INFO**] Using 'UTF-8' encoding to copy filtered resources.
2024-02-14T18:27:42.7598862Z [**INFO**] Copying 34 resources
2024-02-14T18:27:42.7606198Z [**INFO**] 
2024-02-14T18:27:42.7608826Z [**INFO**] --- maven-compiler-plugin:3.11.0:testCompile (default-testCompile) @ classic-templates ---
2024-02-14T18:27:42.7787768Z [**INFO**] Changes detected - recompiling the module! :dependency
2024-02-14T18:27:42.8088263Z [**INFO**] Compiling 136 source files with javac [debug target 11] to target/test-classes
2024-02-14T18:27:47.8184307Z [**WARNING**] /home/fedja/Desktop/my_runner/actions-runner/_work/PrivateCITest/PrivateCITest/v1/src/test/java/com/google/cloud/teleport/spanner/ddl/RandomDdlGenerator.java:[107,25] [AutoValueMutable] An @AutoValue property that is a primitive array returns the original array, which can therefore be modified by the caller. If this is OK, you can suppress this warning with @SuppressWarnings("mutable"). Otherwise, you should replace the property with an immutable type, perhaps a simple wrapper around the original array.
2024-02-14T18:27:47.8194687Z [**WARNING**] /home/fedja/Desktop/my_runner/actions-runner/_work/PrivateCITest/PrivateCITest/v1/src/test/java/com/google/cloud/teleport/templates/SpannerToTextIT.java:[114,57] non-varargs call of varargs method with inexact argument type for last parameter;
2024-02-14T18:27:47.8198629Z   cast to java.lang.Object for a varargs call
2024-02-14T18:27:47.8200056Z   cast to java.lang.Object[] for a non-varargs call and to suppress this warning
2024-02-14T18:27:47.8204235Z [**WARNING**] /home/fedja/Desktop/my_runner/actions-runner/_work/PrivateCITest/PrivateCITest/v1/src/test/java/com/google/cloud/teleport/templates/SpannerToTextIT.java:[173,57] non-varargs call of varargs method with inexact argument type for last parameter;
2024-02-14T18:27:47.8207777Z   cast to java.lang.Object for a varargs call
2024-02-14T18:27:47.8209375Z   cast to java.lang.Object[] for a non-varargs call and to suppress this warning
2024-02-14T18:27:47.8211929Z [**WARNING**] Implicitly compiled files were not subject to annotation processing.
2024-02-14T18:27:47.8214570Z   Use -proc:none to disable annotation processing or -implicit to specify a policy for implicit compilation.
2024-02-14T18:27:47.8219451Z [**INFO**] /home/fedja/Desktop/my_runner/actions-runner/_work/PrivateCITest/PrivateCITest/v1/src/test/java/com/google/cloud/teleport/util/DatadogApiKeyNestedValueProviderTest.java: Some input files use or override a deprecated API.
2024-02-14T18:27:47.8225082Z [**INFO**] /home/fedja/Desktop/my_runner/actions-runner/_work/PrivateCITest/PrivateCITest/v1/src/test/java/com/google/cloud/teleport/util/DatadogApiKeyNestedValueProviderTest.java: Recompile with -Xlint:deprecation for details.
2024-02-14T18:27:47.8230510Z [**INFO**] /home/fedja/Desktop/my_runner/actions-runner/_work/PrivateCITest/PrivateCITest/v1/src/test/java/com/google/cloud/teleport/spanner/ExportTransformTest.java: Some input files use unchecked or unsafe operations.
2024-02-14T18:27:47.8236221Z [**INFO**] /home/fedja/Desktop/my_runner/actions-runner/_work/PrivateCITest/PrivateCITest/v1/src/test/java/com/google/cloud/teleport/spanner/ExportTransformTest.java: Recompile with -Xlint:unchecked for details.
2024-02-14T18:27:47.8238860Z [**INFO**] 
2024-02-14T18:27:47.8239981Z [**INFO**] --- maven-checkstyle-plugin:3.2.1:check (default) @ classic-templates ---
2024-02-14T18:27:48.0263900Z [**INFO**] 
2024-02-14T18:27:48.0266253Z [**INFO**] --- maven-surefire-plugin:2.21.0:test (default-test) @ classic-templates ---
2024-02-14T18:27:48.1134211Z [**INFO**] Surefire report directory: /home/fedja/Desktop/my_runner/actions-runner/_work/PrivateCITest/PrivateCITest/v1/target/surefire-reports
2024-02-14T18:27:48.1164930Z [**INFO**] parallel='none', perCoreThreadCount=true, threadCount=0, useUnlimitedThreads=false, threadCountSuites=0, threadCountClasses=0, threadCountMethods=0, parallelOptimized=true
2024-02-14T18:27:48.2449817Z [**INFO**] 
2024-02-14T18:27:48.2450802Z [**INFO**] -------------------------------------------------------
2024-02-14T18:27:48.2451622Z [**INFO**]  T E S T S
2024-02-14T18:27:48.2452350Z [**INFO**] -------------------------------------------------------
2024-02-14T18:27:49.4600672Z [**INFO**] Running com.google.cloud.teleport.templates.common.DatastoreConvertersTest
2024-02-14T18:27:49.4779388Z SLF4J: Class path contains multiple SLF4J bindings.
2024-02-14T18:27:49.4782627Z SLF4J: Found binding in [jar:file:/home/fedja/.m2/repository/org/slf4j/slf4j-reload4j/1.7.36/slf4j-reload4j-1.7.36.jar!/org/slf4j/impl/StaticLoggerBinder.class]
2024-02-14T18:27:49.4785374Z SLF4J: Found binding in [jar:file:/home/fedja/.m2/repository/org/slf4j/slf4j-log4j12/1.7.25/slf4j-log4j12-1.7.25.jar!/org/slf4j/impl/StaticLoggerBinder.class]
2024-02-14T18:27:49.4788905Z SLF4J: Found binding in [jar:file:/home/fedja/.m2/repository/org/slf4j/slf4j-simple/1.7.36/slf4j-simple-1.7.36.jar!/org/slf4j/impl/StaticLoggerBinder.class]
2024-02-14T18:27:49.4790820Z SLF4J: See http://www.slf4j.org/codes.html#multiple_bindings for an explanation.
2024-02-14T18:27:49.6481730Z SLF4J: Actual binding is of type [org.slf4j.impl.Reload4jLoggerFactory]
2024-02-14T18:27:55.3850340Z [**INFO**] Tests run: 1, Failures: 0, Errors: 0, Skipped: 0, Time elapsed: 5.898 s - in com.google.cloud.teleport.templates.common.DatastoreConvertersTest
2024-02-14T18:27:55.9089142Z [**INFO**] 
2024-02-14T18:27:55.9090179Z [**INFO**] Results:
2024-02-14T18:27:55.9092265Z [**INFO**] 
2024-02-14T18:27:55.9093815Z [**INFO**] Tests run: 1, Failures: 0, Errors: 0, Skipped: 0
2024-02-14T18:27:55.9095294Z [**INFO**] 
2024-02-14T18:27:55.9177971Z [**INFO**] ------------------------------------------------------------------------
2024-02-14T18:27:55.9179857Z [**INFO**] BUILD SUCCESS
2024-02-14T18:27:55.9181503Z [**INFO**] ------------------------------------------------------------------------
2024-02-14T18:27:55.9184499Z [**INFO**] Total time:  43.717 s
2024-02-14T18:27:55.9189770Z [**INFO**] Finished at: 2024-02-14T19:27:55+01:00
2024-02-14T18:27:55.9192695Z [**INFO**] ------------------------------------------------------------------------
2024-02-14T18:27:56.0395252Z 2024/02/14 19:27:56 Test check without NonDex completed successfully!
