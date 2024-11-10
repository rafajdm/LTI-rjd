import React from 'react';
import CandidateProfileForm from '../components/CandidateProfileForm';
import Layout from '../components/Layout';

const HomePage = () => {
  return (
    <Layout>
      <div className="max-w-2xl mx-auto">
        <h1 className="text-2xl font-bold text-gray-900 mb-6">Candidate Profile Form</h1>
        <CandidateProfileForm />
      </div>
    </Layout>
  );
};

export default HomePage;