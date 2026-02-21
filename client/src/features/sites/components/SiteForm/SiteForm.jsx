import {useState} from 'react';
import './SiteForm.css';

const SiteForm = (props) => {
  const [url, setUrl] = useState('');
  const [name, setName] = useState('');
  const [isSubmitting, setIsSubmitting] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (!name || !url) return;

    setIsSubmitting(true);
    try {
      await props.onAdd({name, url});
      setName('');
      setUrl('');
    } catch (err) {
      alert('Error adding site form');
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <div className="site-form-container">
      <h3 style={{marginTop: 0, marginBottom: '16px'}}>Add new service</h3>
      <form className="site-form" onSubmit={handleSubmit}>
        <input
          type="text"
          placeholder="Name (Google)"
          value={name}
          onChange={(e) => setName(e.target.value)}
          required
        />
        <input
          type="url"
          placeholder="URL (https://...)"
          value={url}
          onChange={(e) => setUrl(e.target.value)}
          required
        />
        <button
          type="submit"
          className="submit-btn"
          disabled={isSubmitting}
        >
          {isSubmitting ? 'Submitting...' : 'Add'}
        </button>
      </form>
    </div>
  );
};

export default SiteForm;