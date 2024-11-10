export const createCandidateProfile = async (formData) => {
  const response = await fetch('/api/candidates', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(formData)
  });

  if (!response.ok) {
    throw new Error('Failed to create candidate profile');
  }

  return response.json();
};