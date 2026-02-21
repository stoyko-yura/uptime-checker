import React from 'react';
import SiteCard from '../SiteCard/SiteCard.jsx';
import './SiteList.css';

const SiteList = (props) => {
  if (props.loading) return <p>Loading...</p>;

  if (props.sites.length === 0) {
    return (
      <div className="empty-state">
        <p>You already have no sites.</p>
        <small>Add your first URL, to start monitoring.</small>
      </div>
    );
  }

  return (
    <div className="site-list">
      {props.sites.map((site) => (
        <SiteCard
          key={site.id}
          site={site}
          onDelete={props.onDelete}
        />
      ))}
    </div>
  );
};

export default SiteList;