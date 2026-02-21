import React from 'react';
import './SiteCard.css'

const siteCard = ({site, onDelete}) => {
  const isOnline = site.last_status === 200;

  return (
    <div className="site-card">
      <div className="site-info">
        <span className="site-name">{site.name || 'No named'}</span>
        <a href={site.url} target="_blank" rel="noreferrer" className="site-url">{site.url}</a>
      </div>

      <div style={{display: 'flex', alignItems: 'center', gap: '20px'}}>
        <div className={`status-badge ${isOnline ? 'status-online' : 'status-offline'}`}>
          <div className={`status-dot ${isOnline ? 'dot-online' : 'dot-offline'}`}/>
          {isOnline ? 'Online' : 'Offline'}
        </div>

        <button className="delete-btn" onClick={() => onDelete(site.id)}>
          Delete
        </button>
      </div>
    </div>
  );
};

export default siteCard;