import React, {SyntheticEvent, useEffect, useState} from 'react';
import TopNavbar from "../../components/TopNavbar";
import AppLayout from "../../components/AppLayout";
import {Container, Image, Form, Button, Row} from "react-bootstrap";
import {nanoid} from "nanoid";
import autolab_logo from '../../images/autolab_logo.png';
import {getBackendApiUrl, getFrontendUrl} from "../../utils/url";
import {NavigateFunction, useNavigate} from "react-router-dom";

const axios = require('axios').default;

const RedirectPage = ({pageLink, navigate}: { pageLink: string; navigate: NavigateFunction }) => {
    const handleSubmit = (e: SyntheticEvent) => {
        e.preventDefault();
        const $authCode = document.getElementById("AutolabAuthCode") as HTMLInputElement;
        const $stateCode = localStorage.getItem('authStateValue');
        console.log($authCode.value);
        const url = "/oauth-callback?code=" + $authCode.value + "&state=" + $stateCode;
        navigate(url, {replace: true});
    }

    return (
            <AppLayout>
                <Row className="mb-3">
                    <TopNavbar/>
                </Row>
                <div className="col col-md-8 mx-auto">
                    <h2>Bind with Autolab</h2>
                    <Container className="mt-3 mb-5">
                        <Container className="text-start">
                            <p>Click the banner to visit Autolab authorization page.</p>
                            <p>After the authorization, there will be an authorization code. Paste the code here to
                                finish binding.</p>
                        </Container>
                        <a target="_blank" href={pageLink} rel="noreferrer">
                            <Image src={autolab_logo} alt="Autolab Logo">
                            </Image>
                        </a>
                    </Container>
                    <Form className="mb-3" onSubmit={e => {
                        handleSubmit(e)
                    }}>
                        <Form.Group controlId="AutolabAuthCode">
                            <Form.Control type="text" placeholder="Enter Auth Code"/>
                        </Form.Group>
                        <Button className="mt-3" variant="primary" type="submit">
                            Submit
                        </Button>
                    </Form>
                </div>
            </AppLayout>
    )
}

function AuthRedirect() {

    const [page, setPage] = useState(<div><h1>Redirecting...</h1></div>);
    const navigate = useNavigate();

    const renderPage = async (navigate: NavigateFunction) => {
        // Autolab link URL
        let autolabLink: string = process.env.REACT_APP_AUTOLAB_LOCATION + "/oauth/authorize";

        // Backend Authentication API URL, detect whether I need auth (returns Client ID & scope), or I already had token and good to go.
        const backendAuthUrl = getBackendApiUrl("/auth/info");
        const authInfo = await axios.get(backendAuthUrl);

        // If I am good to go, go to dashboard
        if (authInfo.type === 0) {
            window.location.replace(getFrontendUrl('/dashboard'));
        }

        // Otherwise I need authentication. Keep going through OAuth2 process.

        // Client ID
        let clientId: string = authInfo.data.data.clientId;

        if (process.env.REACT_APP_CLIENT_ID) {
            clientId = process.env.REACT_APP_CLIENT_ID as string;
        }
        autolabLink += `?response_type=code&client_id=${clientId}`;

        // Scope
        //const scope = authInfo.data.scope || "user_info user_courses user_scores user_submit instructor_all admin_all";
        const scope = "user_info user_courses user_scores user_submit instructor_all admin_all";
        autolabLink += `&scope=` + encodeURIComponent(scope);

        // State
        const stateValue = nanoid();
        localStorage.setItem('authStateValue', stateValue);
        autolabLink += `&state=` + stateValue;

        // RedirectURI: if local, display the page
        if (process.env.REACT_APP_REDIRECT_URI === 'urn:ietf:wg:oauth:2.0:oob') {
            autolabLink += '&redirect_uri=' + encodeURIComponent(process.env.REACT_APP_REDIRECT_URI);
            setPage(<RedirectPage pageLink={autolabLink} navigate={navigate}/>);
        } else {
            // Otherwise, set up redirectURI and redirect to Autolab.
            let redirectUri = getFrontendUrl('/oauth-callback');

            redirectUri = encodeURIComponent(process.env.REACT_APP_REDIRECT_URI || redirectUri);
            autolabLink += '&redirect_uri=' + redirectUri;
            window.location.replace(autolabLink);
        }
    }

    useEffect(() => {
        renderPage(navigate)
            .catch(error => console.log("Bad Rendering"));
    }, [navigate])

    return <>{page}</>
}

export default AuthRedirect;