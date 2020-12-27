# Drone Settings Plugins

[![Build Status](https://cloud.drone.io/api/badges/khmarbaise/drone-settings/status.svg)](https://cloud.drone.io/khmarbaise/drone-settings)

# Status

The current state of development is `Prototype`. Not usable at the moment.

# Overview

This represents a Drone.io plugin to handle `settings.xml` for Maven based builds.

## Basic Idea

The basic idea is to have support for a `settings.xml` file within a drone build pipeline.

This will prevent to have each project checked in a `settings.xml` file.

Usually the `settings.xml` contains the configuration from where to consume the needed artifacts. In
a corporate environments this is usually configured to use an internal repository manager.

An example for such `settings.xml` looks like the following:

```xml

<settings xmlns="http://maven.apache.org/SETTINGS/1.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://maven.apache.org/SETTINGS/1.0.0 http://maven.apache.org/xsd/settings-1.0.0.xsd">
  <mirrors>
    <mirror>
      <!--
       ! This sends everything else to /maven-public
      -->
      <id>nexus</id>
      <mirrorOf>*</mirrorOf>
      <url>http://192.168.0.110:8081/repository/maven-public</url>
    </mirror>
  </mirrors>
  <profiles>
    <profile>
      <id>nexus</id>
      <repositories>
        <repository>
          <id>central</id>
          <url>http://central</url>
          <releases>
            <enabled>true</enabled>
            <updatePolicy>never</updatePolicy>
            <checksumPolicy>fail</checksumPolicy>
          </releases>
          <snapshots>
            <enabled>true</enabled>
            <updatePolicy>always</updatePolicy>
            <checksumPolicy>fail</checksumPolicy>
          </snapshots>
        </repository>
      </repositories>
      <pluginRepositories>
        <pluginRepository>
          <id>central</id>
          <url>http://central</url>
          <releases>
            <enabled>true</enabled>
            <updatePolicy>never</updatePolicy>
            <checksumPolicy>fail</checksumPolicy>
          </releases>
          <snapshots>
            <enabled>true</enabled>
            <updatePolicy>always</updatePolicy>
            <checksumPolicy>fail</checksumPolicy>
          </snapshots>
        </pluginRepository>
      </pluginRepositories>
      <properties>
        <distributionManagementRelease>http://192.168.0.110:8081/repository/maven-releases/</distributionManagementRelease>
        <distributionManagementSnapshots>http://192.168.0.110:8081/repository/maven-snapshots/</distributionManagementSnapshots>
        <distributionManagementSite>http://192.168.0.110:8081/repository/maven-sites/</distributionManagementSite>
      </properties>
    </profile>
  </profiles>
  <activeProfiles>
    <!--make the profile active all the time -->
    <activeProfile>nexus</activeProfile>
  </activeProfiles>
</settings>
``` 

This is simply using a defined file at a singe location.

The issue related to Drone.io is that all builds can define their own usage of plugins in
their `.drone.yml`
file. This would mean not to save a single pice of code.

This should support using of credentials for access to the repository manager which should not been
written directly into the `settings.xml`.

# TODO

Using a template for `settings.xml` and fill in the needed fields.

Added server entries with support for keys:

```xml
<settings xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/settings-1.0.0.xsd">

  <mirrors>
    <mirror>
      <!--This sends everything else to /public -->
      <id>nexus</id>
      <mirrorOf>*</mirrorOf>
      <url>http://localhost:8081/nexus/content/groups/public</url>
    </mirror>
  </mirrors>
  <profiles>
    <profile>
      <id>nexus</id>
      <!--Enable snapshots for the built in central repo to direct -->
      <!--all requests to nexus via the mirror -->
      <repositories>
        <repository>
          <id>central</id>
          <url>http://central</url>
          <releases>
            <enabled>true</enabled>
            <checksumPolicy>fail</checksumPolicy>
          </releases>
          <snapshots>
            <enabled>true</enabled>
            <checksumPolicy>fail</checksumPolicy>
          </snapshots>
        </repository>
      </repositories>
      <pluginRepositories>
        <pluginRepository>
          <id>central</id>
          <url>http://central</url>
          <releases>
            <enabled>true</enabled>
            <checksumPolicy>fail</checksumPolicy>
          </releases>
          <snapshots>
            <enabled>true</enabled>
            <checksumPolicy>fail</checksumPolicy>
          </snapshots>
        </pluginRepository>
      </pluginRepositories>
    </profile>
    <profile>
      <id>apache-release</id>
      <properties>
        <gpg.keyname>160788A2</gpg.keyname>
      </properties>
    </profile>
  </profiles>
  <activeProfiles>
    <!--make the profile active all the time -->
    <activeProfile>nexus</activeProfile>
  </activeProfiles>

  <servers>
    <server>
      <id>mavencasts.com</id>
      <username>root</username>
      <privateKey>/Users/kama/keys/rsa-h5593.serverkompetenz.net.ppk</privateKey>
    </server>
    <!-- To stage a website of some part of Maven -->
    <server>
      <id>stagingSite</id> <!-- must match hard-coded repository identifier in site:stage-deploy -->
      <username>khmarbaise</username>
      <password>{fsPqBHSSFQIH+velFoxYZiUBs7uuN25/KP8o61kNS00=}</password>
      <filePermissions>664</filePermissions>
      <directoryPermissions>775</directoryPermissions>
    </server>

    <server>
      <id>site</id>
      <username>admin</username>
      <password>admin123</password>
    </server>
    <server>
      <id>releases</id>
      <username>admin</username>
      <password>admin123</password>
    </server>
    <server>
      <id>snapshots</id>
      <username>admin</username>
      <password>admin123</password>
    </server>

  </servers>

</settings>

```


The idea is to have within a drone pipeline the options to define a 
user defined `settings.xml` which can be enhanced with entries for servers etc.
to make deployment to a repository manager easy possible.