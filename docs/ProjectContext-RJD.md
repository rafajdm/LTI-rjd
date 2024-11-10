# **LTI Applicant Tracking System (ATS) Documentation**

## **Table of Contents**

- [**LTI Applicant Tracking System (ATS) Documentation**](#lti-applicant-tracking-system-ats-documentation)
  - [**Table of Contents**](#table-of-contents)
  - [**Software Description**](#software-description)
    - [**Added Value and Advantages**](#added-value-and-advantages)
    - [**Lean Canvas**](#lean-canvas)
  - [**Main Use Cases**](#main-use-cases)
    - [**Use Case 1: AI-Powered Candidate Screening**](#use-case-1-ai-powered-candidate-screening)
    - [**Use Case 2: Collaborative Interview Scheduling**](#use-case-2-collaborative-interview-scheduling)
    - [**Use Case 3: Automated Candidate Communication**](#use-case-3-automated-candidate-communication)
  - [**Data Model**](#data-model)
    - [**Class Diagram for Data Model**](#class-diagram-for-data-model)
  - [**C4 Model Diagrams**](#c4-model-diagrams)
    - [**Level 1: System Context Diagram**](#level-1-system-context-diagram)
    - [**Level 2: Container Diagram**](#level-2-container-diagram)
    - [**Level 3: Component Diagram for Use Case 3**](#level-3-component-diagram-for-use-case-3)
    - [**Level 4: Code Structure Diagram for Use Case 3**](#level-4-code-structure-diagram-for-use-case-3)

---

## **Software Description**

LTI is an innovative Applicant Tracking System (ATS) designed specifically for medium-sized businesses. It leverages advanced AI data analysis to streamline the recruitment process, enhance HR efficiency, and foster real-time collaboration between recruiters and hiring managers. By automating routine tasks and providing insightful analytics, LTI empowers organizations to make faster and more informed hiring decisions.

### **Added Value and Advantages**

- **AI-Powered Applicant Screening**: Utilizes machine learning algorithms to analyze resumes and rank candidates based on fit, significantly reducing manual screening time.
- **Real-Time Collaboration Tools**: Integrated platforms for communication between recruiters and hiring managers.
- **Automation of Routine Tasks**: Frees up HR professionals by automating scheduling and notifications.
- **User-Friendly Interface**: Quick adoption with minimal training.
- **Advanced Analytics**: Actionable insights via dashboards.
- **Scalability**: Easily integrates with existing HR systems and scales with the business.

### **Lean Canvas**

```mermaid
flowchart TD
    Problem[Problem: Manual screening, collaboration challenges, inefficiencies]
    Solution[Solution: AI screening, real-time collaboration, task automation]
    KeyMetrics[Key Metrics: Time to fill positions, increase productivity, engagement]
    UniqueValue[Unique Value: Accelerate hiring with AI-driven insights]
    UnfairAdvantage[Unfair Advantage: Proprietary AI algorithms, seamless integrations]
    CustomerSegments[Customer Segments: HR departments of medium-sized businesses]
    Channels[Channels: Sales, partnerships, online marketing]
    CostStructure[Cost Structure: Development, AI infrastructure, marketing]
    RevenueStreams[Revenue Streams: Subscriptions, premium features, customization]

    Problem --> Solution
    Solution --> KeyMetrics
    Solution --> UniqueValue
    UniqueValue --> CustomerSegments
    UniqueValue --> UnfairAdvantage
    CustomerSegments --> Channels
    UniqueValue --> CostStructure
    UniqueValue --> RevenueStreams
```

---

## **Main Use Cases**

### **Use Case 1: AI-Powered Candidate Screening**

```plantuml
@startuml
actor Candidate
actor Recruiter
usecase SubmitApplication as UC1
usecase ParseAndRank as UC2
usecase ReviewCandidates as UC3
usecase ShortlistCandidates as UC4

Candidate --> UC1
UC1 --> UC2
UC2 --> Recruiter : Notify about Rankings
Recruiter --> UC3
Recruiter --> UC4
@enduml
```

### **Use Case 2: Collaborative Interview Scheduling**

```plantuml
@startuml
actor Recruiter
actor HiringManager
actor Candidate
usecase InitiateScheduling as UC1
usecase CheckAvailability as UC2
usecase SendInvitation as UC3
usecase ConfirmSchedule as UC4

Recruiter --> UC1
UC1 --> UC2
UC2 --> HiringManager : Confirm Availability
UC2 --> Candidate : Send Invitation
HiringManager --> UC4
Candidate --> UC4
@enduml
```

### **Use Case 3: Automated Candidate Communication**

```plantuml
@startuml
actor Candidate
actor Recruiter
usecase SendAcknowledgment as UC1
usecase SendStatusUpdate as UC2
usecase SendInterviewInstructions as UC3
usecase SendOffer as UC4
usecase RequestFeedback as UC5

UC1 --> Candidate
UC2 --> Candidate
UC3 --> Candidate
UC4 --> Candidate
UC5 --> Candidate
Recruiter --> UC2 : Optionally Customize Messages
@enduml
```

---

## **Data Model**

### **Class Diagram for Data Model**

```mermaid
classDiagram
    class User {
        +UUID userID
        +String role
        +String firstName
        +String lastName
        +String email
    }
    class Candidate {
        +UUID candidateID
        +String firstName
        +String lastName
        +UUID resumeID
        +String email
    }
    class JobPosting {
        +UUID jobID
        +String title
        +String description
    }
    class Application {
        +UUID applicationID
        +UUID candidateID
        +UUID jobID
        +Date appliedAt
    }
    class Interview {
        +UUID interviewID
        +UUID applicationID
        +Date scheduledAt
        +String status
    }
    class AIAnalysis {
        +UUID analysisID
        +UUID applicationID
        +Float score
        +String insights
    }
    class Feedback {
        +UUID feedbackID
        +UUID interviewID
        +UUID reviewerID
        +String comments
        +Int rating
        +Date createdAt
    }

    User "1" --> "*" JobPosting : creates
    User "1" --> "*" Interview : schedules
    User "1" --> "*" Feedback : provides
    Candidate "1" --> "*" Application : submits
    Candidate "1" --> "1" Resume : owns
    Application "1" --> "1" AIAnalysis : triggers
    Application "1" --> "*" Interview : has
    Interview "1" --> "1" Feedback : receives
```

---

## **C4 Model Diagrams**

### **Level 1: System Context Diagram**

This diagram provides a high-level overview of how external users and systems interact with the ATS.

```mermaid
flowchart TD
    Recruiters[Recruiters] --> ATS
    HiringManagers[Hiring Managers] --> ATS
    Candidates[Candidates] --> ATS
    ATS --> JobBoards[Job Boards]
    ATS --> EmailServers[Email Servers]
    ATS --> CalendarSystems[Calendar Systems]
```

### **Level 2: Container Diagram**

This diagram shows the containers within the ATS and how they communicate.

```mermaid
flowchart TD
    subgraph ATS
        WebApp[Web Application]
        API[API Server]
        AI[AI Services]
        Database[(Relational Database)]
    end
    WebApp --> API
    API --> AI
    API --> Database
```

### **Level 3: Component Diagram for Use Case 3**

```mermaid
flowchart TD
    subgraph CommunicationService
        Scheduler[Message Scheduler]
        TemplateEngine[Template Engine]
        NotificationService[Notification Service]
        EmailService[Email Service]
    end
    WebApp --> CommunicationService
    CommunicationService --> Scheduler
    Scheduler --> TemplateEngine
    TemplateEngine --> NotificationService
    NotificationService --> EmailService
```

### **Level 4: Code Structure Diagram for Use Case 3**

```mermaid
classDiagram
    class CommunicationService {
        +scheduleMessage()
        +sendNotification()
    }
    class TemplateEngine {
        +loadTemplate()
        +renderTemplate()
    }
    class NotificationService {
        +sendEmail()
        +sendSMS()
    }
    class EmailService {
        +send(email: Email)
    }
    CommunicationService --> TemplateEngine
    CommunicationService --> NotificationService
    NotificationService --> EmailService
```