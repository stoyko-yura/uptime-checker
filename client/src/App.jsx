import './App.css';
import {useSites} from './features/sites/hooks/useSites.js';
import SiteForm from './features/sites/components/SiteForm/SiteForm.jsx';
import SiteList from './features/sites/components/SiteList/SiteList.jsx';

function App() {
  const {sites, loading, addSite, deleteSite} = useSites();

  return (
    <div className="app-container">
      <header className="app-header">
        <h1>Uptime</h1>
        <p>Simple & Powerful Monitoring</p>
        <div className="app-stats">
          <span>Total: {sites.length}</span>
          <span>Online: {sites.filter(s => s.last_status === 200).length}</span>
        </div>
      </header>

      <main>
        <SiteForm onAdd={addSite}/>
        <SiteList
          sites={sites}
          loading={loading}
          onDelete={deleteSite}
        />
      </main>
    </div>
  );
}

export default App;
