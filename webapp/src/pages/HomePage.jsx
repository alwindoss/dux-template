import SignIn from "../components/Signin";

const HomePage = (props) => {
    let page;
    if (props.loggedIn) {
        page = <h1>Hello You are logged in</h1>
    } else {
        page = <SignIn />
    }
    return page;
}

export default HomePage;
