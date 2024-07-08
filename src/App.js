import React, { useEffect } from 'react';
import Header from './components/Header';
import About from './components/About';
import Projects from './components/Projects';
import Contact from './components/Contact';
import { setupIntersectionObserver } from './utils/animations';
import './App.css';

function App() {
  useEffect(() => {
    setupIntersectionObserver();
  }, []);

  return (
    <div className="App">
      <Header />
      <main className="main-content">
        <About />
        <Projects />
        <Contact />
      </main>
      <footer className="footer">
        <p>&copy; 2023 Your Name. All rights reserved.</p>
      </footer>
    </div>
  );
}

export default App;