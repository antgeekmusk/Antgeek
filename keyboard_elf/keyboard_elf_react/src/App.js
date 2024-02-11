import React, { useState, useEffect } from 'react';
const tom = 'assets/images/tom.png';
const tomLeg = 'assets/images/tom_leg.png';
const jerry = 'assets/images/jerry.png';
const jerryLeg = 'assets/images/jerry_leg.png';


const Square = ({ speed }) => {
  const style = {
    // width: '100px',
    // height: '100px',
    zIndex: 1,
    transform: `rotate(${speed}deg)`,
    transition: 'transform 0.2s ease',
    objectFit: 'contain',
  };

  return (
    <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>
      <div style={{position: 'relative',display: 'inline-block'}}>
        <img src={tom} alt="Tom"/>
        <img src={tomLeg} alt="Tom Leg" style={{...style, position: 'absolute', bottom: 3, left: 21}}  />
      </div>
      <div style={{position: 'relative',display: 'inline-block'}}>
        <img src={jerry} alt="Jerry"/>
        <img src={jerryLeg} alt="Jerry Leg" style={{...style, position: 'absolute', bottom: 3, left: 10}} />
      </div>
  </div>
  );
};

const App = () => {
  const [speed, setSpeed] = useState(0);
  const [isKeyPressed, setIsKeyPressed] = useState(false);

  useEffect(() => {
    const handleKeyDown = () => {
      if (!isKeyPressed) {
        setIsKeyPressed(true);
        setSpeed(prevSpeed => prevSpeed + 360);
      }
    };

    const handleKeyUp = () => {
      setIsKeyPressed(false);
    };

    document.addEventListener('keydown', handleKeyDown);
    document.addEventListener('keyup', handleKeyUp);

    return () => {
      document.removeEventListener('keydown', handleKeyDown);
      document.removeEventListener('keyup', handleKeyUp); 
    };
  }, [isKeyPressed]);

  return <Square speed={speed} />;
};

export default App;