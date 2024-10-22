import PropTypes from "prop-types";
import { NavLink } from "react-router-dom";

const AppbarLink = (props) => {
	return (
		<NavLink
			key={props.id}
			end={props.end}
			to={props.linkPath}
			className={({ isActive }) =>
				isActive
					? "py-2 px-5 text-xs text-zinc-700 border-r border-white bg-white transition-colors duration-150"
					: "py-2 px-5 text-xs text-zinc-200 hover:text-amber-400 border-r border-white last:border-none transition-colors duration-150"
			}>
			{props.contentText}
		</NavLink>
	);
};

AppbarLink.propTypes = {
	contentText: PropTypes.string,
	end: PropTypes.any,
	id: PropTypes.number,
	linkPath: PropTypes.string
}

export default AppbarLink;
