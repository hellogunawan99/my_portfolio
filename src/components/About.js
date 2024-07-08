import React from 'react';
import './About.css';

const About = () => {
  return (
    <section id="about" className="section about">
      <h2>About Me</h2>
      <div className="about-content">
        <div className="about-text">
          <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>
          <p>Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
        </div>
        <div className="skills">
          <h3>Skills</h3>
          <ul>
            <li>React.js</li>
            <li>Go</li>
            <li>JavaScript</li>
            <li>HTML/CSS</li>
            <li>Node.js</li>
            <li>SQL</li>
          </ul>
        </div>
      </div>
    </section>
  );
};

export default About;