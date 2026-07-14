import { useRef, useEffect } from 'react';


const Tooltip = ({text, children}) => {
    const tooltip = useRef(null);
    const point = useRef(null);
    const component = useRef(null);

    const show = () => {
        tooltip.current.style.display = 'flex';
        const rect = tooltip.current.getBoundingClientRect();
        const right = rect.right + 100;
        const top = rect.top - 100;

        // right
        if(right > innerWidth) {
            tooltip.current.style.right = '0';
            // I'm using tailwind css here because style 
            // doesn't work, which I don't know why
            point.current.classList.add('right-2');
        }
        
        // bottom
        if(top < 0) {
            tooltip.current.style.bottom = `-${rect.height+4}px`;
            point.current.style.cssText = `
                bottom: none;
                top: -4px;
                border: none;
                border-top: 1px solid rgba(255, 255, 255, 0.45);
                border-left: 1px solid rgba(255, 255, 255, 0.45);
            `;
        } else {
            tooltip.current.style.top = `-${rect.height+4}px`;
        }
    }
    
    const hide = () => {
        tooltip.current.style.display = 'none';
    }

    useEffect(() => {
        // I have to save the component(useRef) to a variable(elem), since
        // closure will save this if we save it to a variable as a reference 
        // in the DOM elements, so even this component unmounts when this effect 
        // was setting up, we still have reference to this tooltip(useRef) to 
        // access in the DOM. And this is called captured reference.        
        const elem = component.current;
        elem.addEventListener('mouseenter', show);
        elem.addEventListener('mouseleave', hide);

        return () => {
            elem.removeEventListener('mouseenter', show);
            elem.removeEventListener('mouseleave', hide);
        }
    }, []);

    return (
        <div ref={component} className="relative cursor-pointer">
            <div ref={tooltip} className="absolute hidden bg-black p-1 pb-1.5 rounded-sm border border-white/45 z-50">
                <span className="text-nowrap text-white leading-none">
                    {text}
                </span>
                <div ref={point} className="w-2 h-2 bg-black rotate-45 absolute -bottom-1 border-b border-r border-white/45"></div>
            </div>
            {children}
        </div>
    );
}

export default Tooltip;