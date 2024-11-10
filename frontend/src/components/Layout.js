import React from 'react';

const Layout = ({ children }) => {
  return (
    <div className="flex h-screen bg-gray-100">
      {/* Sidebar */}
      <div className="w-64 bg-white shadow-lg">
        <div className="p-4">
          <h1 className="text-2xl font-semibold text-gray-800">LTI ATS</h1>
        </div>
        <nav className="mt-4">
          <a href="/" className="block px-4 py-2 text-gray-600 hover:bg-gray-50 hover:text-gray-900">
            Dashboard
          </a>
          <a href="/candidates" className="block px-4 py-2 text-gray-600 hover:bg-gray-50 hover:text-gray-900">
            Candidates
          </a>
          <a href="/jobs" className="block px-4 py-2 text-gray-600 hover:bg-gray-50 hover:text-gray-900">
            Jobs
          </a>
          <a href="/interviews" className="block px-4 py-2 text-gray-600 hover:bg-gray-50 hover:text-gray-900">
            Interviews
          </a>
        </nav>
      </div>

      {/* Main content */}
      <div className="flex-1 overflow-auto">
        <div className="p-8">
          {children}
        </div>
      </div>
    </div>
  );
};

export default Layout;