# CD Flow

Simple framework and binary to emit Cloud Events related to Continuous Delivery. 

The idea behind this framework and command-line interface is to empower you  emitting events related to a Continuous Delivery flow to measure performance and gain visibility of your processes. 

## Supported Events

## CDF Issue Events

- Base
  - Title
  - Description
  - URL
  - Author
- Lifecycle Events
  - **CDF.Issue.Created**
  - **CDF.Issue.CommentAdded**
  - **CDF.Issue.Closed**


## CDF Repository Events
- Base
  - Name
  - Url 
- PR Events
  - **CDF.Repository.PR.Created**
    - PR Url
    - PR Title
    - Issue
  - **CDF.Repository.PR.Merged**
  - **CDF.Repository.PR.CommentAdded**
- Main Events
  - **CDF.Repository.Main.Changed**
- Tags Events
    - **CDF.Repository.Tag.Created**

## CDF Artifact Events
- Base
  - Name
  - Version
  - Source
- Lifecycle Events:
  - **CDF.Artifact.TestStarted**
  - **CDF.Artifact.TestEnded**  
    - TestPassed
  - **CDF.Artifact.Built**
    - SHA 
  - **CDF.Artifact.VersionUpdated**
    - OldVersion
    - NewVersion
  - **CDF.Artifact.Released**
    - URL
    - Kind
      - Library
      - Service
## CDF Container Events
- Base
  - Name
  - Organization
  - Tag
  - Repository
- Lifecycle Events
  - **CDF.Container.Built**
  - **CDF.Container.Released**
  
## CDF Service Events
- Base
  - Name
  - Url
  - Version
  - Environment URL
- Lifecycle Events
  - **CDF.Service.Deployed**
  - **CDF.Service.Upgraded**

## CDF Environment Events
- Base
  - URL
  - Repository
  - Name
- Lifecycle Events
  - **CDF.Environment.Created**
  - **CDF.Environment.Updated**
  - **CDF.Environment.ServicePromoted**


## Visualization
(TBD)

