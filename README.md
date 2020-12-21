# CD Flow

Simple framework and binary to emit Cloud Events related to Continuous Delivery. 

The idea behind this framework and command-line interface is to empower you  emitting events related to a Continuous Delivery flow to measure performance and gain visibility of your processes. 

## Supported Events

Use `cdf --help`

The following events are currently supported:
- [Issue Events]()
- [Repository Events]()
- [Artifact Events]()
- [Container Events]()
- [Service Events]()
- [Environment Events]()


## CDF Issue Events

Use `cdf issue --help`

- Base Properties
  - Id (required)
  - Repository (required)
  - Title 
  - Author 
  
### **CDF.Issue.Created** Event

Example: `./cdf issue created --id 7 --repository my-project --title "new issue" --author salaboy`


### **CDF.Issue.Updated** Event


Example: `./cdf issue updated --id 7 --repository my-project`

### **CDF.Issue.CommentAdded** Event

Example: `./cdf issue commented --id 7 --repository my-project --comment "new comment"`

### **CDF.Issue.Closed** Event

Example: `./cdf issue closed --id 7 --repository my-project`


## CDF Repository Events

Use `cdf repo --help`

- Base Properties
  - Name
  - Url 
  

### Pull Request Events

Use `cdf repo pr --help`

- Base Properties    
    - PR Id (required)
    - PR Repository (required)
    - Issue Id
    - PR Title


#### **CDF.Repository.PR.Created** Event

Example: `./cdf repo pr created --id 7 --repository my-project`

#### **CDF.Repository.PR.Merged** Event

Example: `./cdf repo pr merged --id 7 --repository my-project`

#### **CDF.Repository.PR.CommentAdded** Event

Example: `./cdf repo pr comment --id 7 --repository my-project`

### Main Events

#### **CDF.Repository.Main.Changed** Event

### Tags Events


#### **CDF.Repository.Tag.Created** Event

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

