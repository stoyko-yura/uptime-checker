import {useEffect, useState} from 'react';
import api from '../../../api/axios.js';

export const useSites = () => {
  const [sites, setSites] = useState([]);
  const [loading, setLoading] = useState(true);

  const fetchSites = async () => {
    try {
      const response = await api.get('/sites');
      setSites(response.data);
    } catch (err) {
      console.error('Fetch error', err);
    } finally {
      setLoading(false);
    }
  };

  const addSite = async (siteData) => {
    const response = await api.post('/sites', siteData);
    setSites(prev => [...prev, response.data]);
  };

  const deleteSite = async (id) => {
    await api.delete(`/sites/${id}`);
    setSites(prev => prev.filter(s => s.id !== id));
  };

  useEffect(() => {
    fetchSites();
    const interval = setInterval(fetchSites, 10000);
    return () => clearInterval(interval);
  }, []);

  return {sites, loading, addSite, deleteSite};
};