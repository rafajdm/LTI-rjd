-- migrations/20231010120000_create_candidate_profile_table.up.sql
CREATE TABLE CandidateProfile (
    candidate_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    contact_info JSONB NOT NULL,
    resume_link VARCHAR(255),
    current_position VARCHAR(100),
    location VARCHAR(100)
);