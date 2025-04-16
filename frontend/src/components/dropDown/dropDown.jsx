import { useState } from 'react';

export default function Dropdown({ label, options = [] }) {
  const [open, setOpen] = useState(false);

  const handleOptionClick = (action) => {
    action();
    setOpen(false);
  };

  return (
    <div className="relative inline-block text-left">
      <button
        className="text-gray-700 hover:text-blue-600 flex items-center gap-1"
        onClick={() => setOpen(prev => !prev)}
      >
        {label}
      </button>

      {open && (
        <div className="absolute left-1/2 -translate-x-1/2 mt-2 w-48 rounded-md shadow-lg bg-white ring-1 ring-black ring-opacity-5 z-50">
          {options.map(({ label, onClick }, idx) => (
            <button
              key={idx}
              className="block w-full px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 text-left"
              onClick={() => handleOptionClick(onClick)}
            >
              {label}
            </button>
          ))}
        </div>
      )}
    </div>
  );
}
