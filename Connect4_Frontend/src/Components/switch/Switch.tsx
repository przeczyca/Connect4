import './Switch.css'

interface switchProps {
    state: string;
    setState: any;
}

function Switch(props: switchProps) {
    function isChecked() {
        if (props.state === "0") {
            props.setState("2");
        }
        else {
            props.setState("0");
        }
    }

    return (
        <label className="switch">
            <input type="checkbox" onChange={isChecked}/>
            <span className="slider"></span>
        </label>
    );
}

export default Switch;