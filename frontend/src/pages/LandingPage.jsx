import Navbar from '../components/navbar/navbar';
export default function LandingPage() {
  return (
    <div className="min-h-screen flex flex-col bg-white text-gray-900">
      <Navbar />


      <header className="flex flex-col items-center justify-center text-center px-6 py-24 bg-gradient-to-br from-blue-600 to-indigo-700 text-white">
        <h1 className="text-4xl md:text-5xl font-bold mb-4">Run Better Teams</h1>
        <p className="text-lg max-w-xl mb-8">
          Track, manage, and improve your team's agile workflow.
        </p>

        <h2 className="text-2xl font-semibold mb-4">What You Can Do</h2>
        <p className="text-white-600 mb-8">We're just getting started. Soon you'll be able to:</p>
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6 max-w-4xl mx-auto">
          <div className="p-4 bg-white rounded shadow text-gray-500">✅ Start Standups</div>
          <div className="p-4 bg-white rounded shadow text-gray-500">📊 View Issues</div>
          <div className="p-4 bg-white rounded shadow text-gray-500">🛑 End Sessions</div>
        </div>
      </header>

      <footer className="text-center py-6 text-gray-500 border-t">
        © {new Date().getFullYear()} StandupTracker
      </footer>
    </div>
  );
}
