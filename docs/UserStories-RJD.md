
# ATS Project - Plan and Estimation Phase

## User Stories - Introduction
This document lists the user stories for the ATS (Applicant Tracking System) project. Each user story captures a specific feature or functionality, with its acceptance criteria, dependencies, and any special notes. This structure supports an organized development process and provides additional context for each feature.

---

### User Story 1: Capture Basic Candidate Profile Information
- **As a** recruiter, **I want** to create and save a candidate profile with essential information, **so that** I can easily access and review applicant details during the hiring process.
- **Acceptance Criteria**:
  - System captures key candidate information fields: full name, contact details, resume upload, current role, and location.
  - Profile data is stored in the database and can be retrieved for editing.
  - Error message is displayed if required fields are missing.
- **Dependencies**: None
- **Notes**: This story is foundational for all candidate management functionalities.

---

### User Story 2: Update Candidate Application Status
- **As a** recruiter, **I want** to update a candidate's application status, **so that** I can track their progress through different stages of the hiring process.
- **Acceptance Criteria**:
  - System includes predefined application statuses (e.g., Applied, Screening, Interviewing, Offer, Hired, Rejected).
  - Recruiters can update candidate status with a timestamp recorded in the profile history.
  - Invalid status transitions are prevented.
- **Dependencies**: Capture Basic Candidate Profile Information
- **Notes**: Enables recruiters to manage candidate progression through the hiring pipeline.

---

### User Story 3: Search and Filter Candidates
- **As a** recruiter, **I want** to search and filter candidates by various criteria, **so that** I can quickly find specific applicants relevant to current job openings.
- **Acceptance Criteria**:
  - Search bar allows for keyword entry (e.g., candidate name, email).
  - Filters include application status, location, date of application, and skills.
  - If no candidates match, “No Results Found” message appears.
- **Dependencies**: Capture Basic Candidate Profile Information
- **Notes**: Advanced search capabilities improve recruiter efficiency.

---

### User Story 4: AI-Driven Candidate Ranking for Job Applications
- **As a** recruiter, **I want** the ATS to automatically screen and rank candidates based on job fit, **so that** I can prioritize reviewing applicants who closely match the job requirements.
- **Acceptance Criteria**:
  - AI algorithm ranks candidates based on fit factors like experience and skills.
  - Recruiters can view a ranked list with each candidate's fit score displayed.
  - Recruiters can adjust criteria weights for custom ranking.
- **Dependencies**: Capture Basic Candidate Profile Information
- **Notes**: This feature leverages AI for efficient candidate screening, reducing manual review time.

---

### User Story 5: Real-Time Messaging for Candidate Reviews
- **As a** recruiter or hiring manager, **I want** to communicate in real-time with team members about specific candidates, **so that** we can collaborate effectively and make timely hiring decisions.
- **Acceptance Criteria**:
  - Each candidate profile includes a real-time messaging interface for discussions.
  - Notifications alert users of new messages related to specific candidates.
  - Messages are stored and accessible in candidate profile history.
- **Dependencies**: Capture Basic Candidate Profile Information
- **Notes**: Supports collaboration between team members on candidate evaluations.

---

### User Story 6: Automated Interview Scheduling and Notifications
- **As a** recruiter, **I want** the ATS to automatically schedule interviews and send notifications to candidates, **so that** I can save time on administrative tasks and keep candidates informed.
- **Acceptance Criteria**:
  - System allows recruiters to set interview availability for candidates to select time slots.
  - Confirmation and reminder emails are sent to candidates automatically.
  - Changes to the schedule are updated and notified to all parties involved.
- **Dependencies**: Capture Basic Candidate Profile Information, Send Interview Invitation to Candidates
- **Notes**: Reduces administrative workload and improves candidate experience.

---

### User Story 7: Dashboard for Hiring Metrics and Analytics
- **As a** hiring manager, **I want** to view key hiring metrics on a dashboard, **so that** I can monitor recruitment performance and identify areas for improvement.
- **Acceptance Criteria**:
  - Dashboard displays metrics, including time-to-hire and conversion rates.
  - Metrics can be filtered by date range, position, and recruiter.
  - Visual formats include charts and summary tables.
- **Dependencies**: Update Candidate Application Status
- **Notes**: Enables data-driven decision-making by providing insights into recruitment performance.

---

### User Story 8: Generate Hiring Pipeline Report
- **As a** hiring manager, **I want** to generate a report on the current hiring pipeline, **so that** I can assess the progress of recruitment efforts and identify any bottlenecks.
- **Acceptance Criteria**:
  - System generates a report with candidate counts at each hiring stage.
  - Report is viewable on-screen and exportable as a CSV file.
  - Data is presented in both visual (charts) and tabular formats.
- **Dependencies**: Update Candidate Application Status
- **Notes**: Provides detailed visibility into the hiring pipeline for strategic adjustments.

---

### User Story 9: Send Interview Invitation to Candidates
- **As a** recruiter, **I want** to send interview invitations directly to candidates from within the ATS, **so that** I can streamline communication and keep candidates informed of next steps.
- **Acceptance Criteria**:
  - System allows recruiters to send interview invitations including date, time, and additional notes.
  - Invitations are logged in the candidate profile history.
  - Error message appears if sending fails, with retry option.
- **Dependencies**: Capture Basic Candidate Profile Information
- **Notes**: Facilitates interview scheduling directly within the ATS, streamlining communication with candidates.

---

## ATS Project - Backlog Overview - Introduction
This document provides an organized overview of the approved backlog for the ATS (Applicant Tracking System) project. The backlog prioritizes core features and foundational elements based on business value, implementation complexity, and compliance needs. It’s structured in a way that allows progressive development from essential functionalities to advanced features, supporting the project’s MVP goals.

---

## Backlog Structure

### Sprint 1: Core Candidate Management and Security Setup
1. **Create Candidate Profile Table in Database**  
   - **Effort**: 5 (Fibonacci), 8 hours, 3 story points  
   - Description: Setting up the candidate profile data structure in the database.
   
2. **Develop Candidate Profile UI for Data Entry**  
   - **Effort**: 5 (Fibonacci), 8 hours, 3 story points  
   - Description: Creating the UI component for recruiters to enter candidate information.

3. **Implement Secure Data Handling with Encryption**  
   - **Effort**: 8 (Fibonacci), 12 hours, 5 story points  
   - Description: Enabling encryption for data at rest and in transit.

4. **Set Up Role-Based Access Control (RBAC) for Profiles**  
   - **Effort**: 8 (Fibonacci), 10 hours, 5 story points  
   - Description: Implementing RBAC to ensure secure, role-based access to candidate data.

5. **Database Configuration and Setup for Candidate Management**  
   - **Effort**: 3 (Fibonacci), 6 hours, 2 story points  
   - Description: Configuring database connections and optimizing for performance.

6. **Add Update Candidate Status Feature**  
   - **Effort**: 5 (Fibonacci), 8 hours, 3 story points  
   - Description: Adding a feature for recruiters to update candidate status in the hiring pipeline.

7. **Implement Access Logging for Compliance Auditing**  
   - **Effort**: 5 (Fibonacci), 8 hours, 3 story points  
   - Description: Logging access to candidate profiles to meet compliance requirements.

8. **Integrate Logging Framework for System Logs**  
   - **Effort**: 5 (Fibonacci), 8 hours, 3 story points  
   - Description: Integrating a logging framework to capture system logs.

9. **Configure Logging Dashboards and Search Filters**  
   - **Effort**: 5 (Fibonacci), 8 hours, 3 story points  
   - Description: Setting up dashboards for viewing and filtering system logs.

10. **Design Basic Profile Layout and Field Validations**  
    - **Effort**: 3 (Fibonacci), 5 hours, 2 story points  
    - Description: Designing the profile form layout and implementing validations.

11. **Initial Database Backup and Recovery Process Setup**  
    - **Effort**: 5 (Fibonacci), 8 hours, 3 story points  
    - Description: Establishing a backup and recovery plan for candidate data.

12. **Implement API Endpoints for Profile CRUD Operations**  
    - **Effort**: 8 (Fibonacci), 12 hours, 5 story points  
    - Description: Developing endpoints to support CRUD operations on candidate profiles.

---

### Sprint 2: AI Screening and Initial Automation
1. **AI-Driven Candidate Ranking for Job Applications**  
   - **Effort**: 8 (Fibonacci), 10 hours, 5 story points  
   - Description: Using AI to rank candidates based on job fit.

2. **Automated Interview Scheduling and Notifications**  
   - **Effort**: 8 (Fibonacci), 10 hours, 5 story points  
   - Description: Setting up automated scheduling for interviews and notifications.

3. **Candidate Data Backup and Recovery**  
   - **Effort**: 5 (Fibonacci), 8 hours, 3 story points  
   - Description: Establishing a secure backup and recovery system.

---

### Sprint 3: Real-Time Collaboration and Enhanced Security
1. **Real-Time Messaging for Candidate Reviews**  
   - **Effort**: 8 (Fibonacci), 12 hours, 5 story points  
   - Description: Adding real-time messaging to facilitate recruiter communication.

2. **Automated Data Compliance Checks**  
   - **Effort**: 5 (Fibonacci), 8 hours, 3 story points  
   - Description: Automating data checks for compliance.

3. **Basic Notification System**  
   - **Effort**: 5 (Fibonacci), 8 hours, 3 story points  
   - Description: Implementing a notification system to inform users of important updates.

---

### Sprint 4: Analytics and Advanced Reporting
1. **Dashboard for Hiring Metrics and Analytics**  
   - **Effort**: 10 (Fibonacci), 12 hours, 8 story points  
   - Description: Providing a dashboard to display key metrics on the hiring pipeline.

2. **Generate Hiring Pipeline Report**  
   - **Effort**: 8 (Fibonacci), 10 hours, 5 story points  
   - Description: Enabling hiring managers to generate and export pipeline reports.

3. **Advanced Candidate Search and Filtering Options**  
   - **Effort**: 5 (Fibonacci), 8 hours, 3 story points  
   - Description: Adding enhanced search filters for candidate profiles.

---

### Sprint 5: Scalability and Final Enhancements
1. **Integration with External Calendar Systems**  
   - **Effort**: 6 (Fibonacci), 8 hours, 4 story points  
   - Description: Integrating with popular calendars for seamless interview scheduling.

2. **Scalability Testing and Load Balancing**  
   - **Effort**: 5 (Fibonacci), 8 hours, 3 story points  
   - Description: Ensuring the system can handle high data volumes.

3. **Predictive Analytics for Hiring Trends**  
   - **Effort**: 8 (Fibonacci), 10 hours, 5 story points  
   - Description: Adding predictive analytics to forecast hiring needs.

---

## Summary of Total Estimations
- **Total Fibonacci Effort**: 145
- **Total Hours**: 230
- **Total Story Points**: 110

---


## ATS Project - Prioritization Analysis - Introduction
This document outlines the prioritization analysis for the ATS (Applicant Tracking System) project. We applied multiple prioritization frameworks, such as MoSCoW and RICE, to ensure a focus on features that deliver high business and user value, while balancing complexity and dependencies.

---

## MoSCoW Prioritization

The MoSCoW framework categorizes each backlog item based on its necessity and impact, as follows:

- **Must-Have**:
  - Capture Basic Candidate Profile Information
  - Secure Data Handling Setup
  - Update Candidate Application Status
  - Database Setup and Initial Configuration
  - Access Control and Compliance Audit Logging

- **Should-Have**:
  - AI-Driven Candidate Ranking for Job Applications
  - Automated Interview Scheduling and Notifications
  - Basic Notification System
  - Real-Time Messaging for Candidate Reviews
  - Candidate Data Backup and Recovery
  - Automated Data Compliance Checks

- **Could-Have**:
  - Dashboard for Hiring Metrics and Analytics
  - Generate Hiring Pipeline Report
  - Advanced Candidate Search and Filtering Options
  - Integration with External Calendar Systems
  - Scalability Testing and Load Balancing
  - Predictive Analytics for Hiring Trends

---

## RICE Scoring

The RICE framework calculates a score based on Reach, Impact, Confidence, and Effort for each user story. This score helps rank items by potential value and effort required:

| User Story                                      | Reach | Impact | Confidence | Effort | RICE Score |
|-------------------------------------------------|-------|--------|------------|--------|------------|
| Capture Basic Candidate Profile Information     | 1000  | 3      | 0.9        | 5      | 540        |
| Secure Data Handling Setup                      | 1000  | 3      | 0.9        | 5      | 540        |
| Update Candidate Application Status             | 1000  | 2      | 0.8        | 3      | 533        |
| Database Setup and Initial Configuration        | 1000  | 3      | 0.9        | 3      | 900        |
| Access Control and Compliance Audit Logging     | 1000  | 3      | 0.8        | 4      | 600        |
| AI-Driven Candidate Ranking for Job Applications | 800   | 3      | 0.7        | 8      | 210        |
| Automated Interview Scheduling and Notifications | 800   | 3      | 0.8        | 8      | 240        |
| Basic Notification System                       | 800   | 2      | 0.8        | 5      | 256        |
| Real-Time Messaging for Candidate Reviews       | 800   | 2      | 0.7        | 8      | 140        |
| Candidate Data Backup and Recovery              | 1000  | 3      | 0.9        | 5      | 540        |
| Automated Data Compliance Checks                | 1000  | 3      | 0.8        | 4      | 600        |
| Dashboard for Hiring Metrics and Analytics      | 800   | 2      | 0.7        | 10     | 112        |
| Generate Hiring Pipeline Report                 | 800   | 2      | 0.7        | 8      | 140        |
| Advanced Candidate Search and Filtering Options | 800   | 2      | 0.7        | 5      | 224        |
| Integration with External Calendar Systems      | 800   | 2      | 0.7        | 6      | 187        |
| Scalability Testing and Load Balancing          | 1000  | 3      | 0.8        | 5      | 480        |
| Predictive Analytics for Hiring Trends          | 800   | 2      | 0.6        | 8      | 120        |

---

## Summary and Prioritization Decisions

The results from MoSCoW and RICE have been used to structure the backlog to maximize business value and streamline the development process:

1. **Early Focus**: Must-have items and those with high RICE scores, like candidate profile management, secure data handling, and access control, are prioritized to build a strong foundation.
2. **Subsequent Focus**: Should-have items with moderate RICE scores, including AI-driven ranking, automated scheduling, and compliance checks, are planned for implementation in subsequent sprints.
3. **Future Enhancements**: Could-have items with lower scores, such as advanced analytics and external integrations, are set aside for later sprints once core functionality is complete.

This approach balances user and business value with practical considerations of effort and complexity.

---

## ATS Project - Sprint 1 Plan - Overview
This document outlines the Sprint 1 plan for the ATS (Applicant Tracking System) project. It includes approved work tickets, their descriptions, acceptance criteria, and effort estimations in Fibonacci, hours, and story points. Sprint 1 focuses on establishing foundational elements essential for the ATS, with a strong emphasis on candidate profile management, secure data handling, and setting up key backend and security mechanisms.

---

## Work Tickets

### Ticket 1: Create Candidate Profile Table in Database
- **Description**: Create a dedicated table within the database for storing candidate profile information, with necessary fields and constraints for managing candidate data.
- **Acceptance Criteria**:
  - A `CandidateProfile` table exists in the database to store candidate information fields.
  - Required fields include: `candidate_id` (primary key), `name`, `contact_info` (email, phone), `resume_link`, `current_position`, and `location`.
  - Data can be stored and retrieved from the `CandidateProfile` table without errors.
  - Table setup follows defined naming conventions and normalization standards.
- **Effort Estimation**:
  - **Fibonacci**: 5
  - **Hours**: 8 hours
  - **Story Points**: 3 points

---

### Ticket 2: Develop Candidate Profile UI for Data Entry
- **Description**: Create a user interface component for recruiters to input candidate profile information, allowing for efficient data entry.
- **Acceptance Criteria**:
  - A form is available on the UI for entering candidate details: name, contact information (email, phone), resume link, current position, and location.
  - The form includes validation to ensure all required fields are filled correctly.
  - Successful submission adds the candidate data to the `CandidateProfile` table.
  - A confirmation message appears upon successful data entry.
  - Error messages are displayed for invalid or missing data.
- **Effort Estimation**:
  - **Fibonacci**: 5
  - **Hours**: 8 hours
  - **Story Points**: 3 points

---

### Ticket 3: Implement Secure Data Handling with Encryption
- **Description**: Set up encryption mechanisms for data at rest and in transit, ensuring secure handling of candidate data.
- **Acceptance Criteria**:
  - Data at rest within the `CandidateProfile` table is encrypted.
  - HTTPS is enabled for secure data transmission.
  - Encryption protocols meet industry security standards.
  - System prevents unauthorized access to encrypted data.
- **Effort Estimation**:
  - **Fibonacci**: 8
  - **Hours**: 12 hours
  - **Story Points**: 5 points

---

### Ticket 4: Set Up Role-Based Access Control (RBAC) for Profiles
- **Description**: Implement RBAC to restrict access to candidate data based on user roles, ensuring only authorized users can view or modify data.
- **Acceptance Criteria**:
  - Different user roles (e.g., recruiter, hiring manager) are defined with specific access rights.
  - Access control policies ensure authorized access only.
  - Unauthorized access attempts are logged with error notifications.
- **Effort Estimation**:
  - **Fibonacci**: 8
  - **Hours**: 10 hours
  - **Story Points**: 5 points

---

### Ticket 5: Database Configuration and Setup for Candidate Management
- **Description**: Configure the database for the ATS, establishing connections, performance parameters, and enabling CRUD operations.
- **Acceptance Criteria**:
  - Database is configured for CRUD operations on candidate data.
  - Essential performance parameters (e.g., indexing, connection pooling) are set.
  - Security settings restrict unauthorized access.
- **Effort Estimation**:
  - **Fibonacci**: 3
  - **Hours**: 6 hours
  - **Story Points**: 2 points

---

### Ticket 6: Add Update Candidate Status Feature
- **Description**: Develop a feature allowing recruiters to update a candidate’s application status (e.g., Applied, Screening, Interviewing).
- **Acceptance Criteria**:
  - A dropdown menu or selection tool is available for status updates.
  - Status changes are timestamped in the profile history.
  - Confirmation message appears upon update.
- **Effort Estimation**:
  - **Fibonacci**: 5
  - **Hours**: 8 hours
  - **Story Points**: 3 points

---

### Ticket 7: Implement Access Logging for Compliance Auditing
- **Description**: Set up logging to track access to candidate data, meeting compliance requirements.
- **Acceptance Criteria**:
  - Each access or modification of candidate data is logged.
  - Logs include user ID, timestamp, and action.
  - Unauthorized access attempts are also recorded.
- **Effort Estimation**:
  - **Fibonacci**: 5
  - **Hours**: 8 hours
  - **Story Points**: 3 points

---

### Ticket A: Integrate Logging Framework for System Logs
- **Description**: Integrate a logging framework (e.g., Kibana, Splunk) to collect and store system logs, including access and error events.
- **Acceptance Criteria**:
  - Logging framework is installed and connected to the ATS.
  - System logs (e.g., access, errors) are captured within the framework.
- **Effort Estimation**:
  - **Fibonacci**: 5
  - **Hours**: 8 hours
  - **Story Points**: 3 points

---

### Ticket B: Configure Logging Dashboards and Search Filters
- **Description**: Set up dashboards and configure filters within the logging framework for monitoring and compliance.
- **Acceptance Criteria**:
  - Dashboards display key logging metrics (e.g., access attempts).
  - Search filters allow filtering by user ID, timestamp, action.
- **Effort Estimation**:
  - **Fibonacci**: 5
  - **Hours**: 8 hours
  - **Story Points**: 3 points

---

### Ticket 8: Design Basic Profile Layout and Field Validations
- **Description**: Design the profile form layout with field validations to ensure data integrity.
- **Acceptance Criteria**:
  - Form layout for candidate profiles includes fields like name, contact information, and resume link.
  - Validation rules are implemented to ensure data accuracy.
- **Effort Estimation**:
  - **Fibonacci**: 3
  - **Hours**: 5 hours
  - **Story Points**: 2 points

---

### Ticket 9: Initial Database Backup and Recovery Process Setup
- **Description**: Establish a backup and recovery process for the database to prevent data loss.
- **Acceptance Criteria**:
  - Backup process is configured and tested.
  - Recovery procedures allow data restoration from backups.
- **Effort Estimation**:
  - **Fibonacci**: 5
  - **Hours**: 8 hours
  - **Story Points**: 3 points

---

### Ticket 10: Implement API Endpoints for Profile CRUD Operations
- **Description**: Develop API endpoints to support Create, Read, Update, and Delete (CRUD) operations on candidate profiles.
- **Acceptance Criteria**:
  - API endpoints for CRUD operations are implemented.
  - Endpoints include basic validation for data integrity.
  - Testing confirms endpoint functionality.
- **Effort Estimation**:
  - **Fibonacci**: 8
  - **Hours**: 12 hours
  - **Story Points**: 5 points

---

## Summary of Estimations
- **Total Fibonacci Effort**: 65
- **Total Hours**: 109
- **Total Story Points**: 41

---

This document serves as a reference for Sprint 1, detailing each task's requirements and estimations. Approved for planning and execution.