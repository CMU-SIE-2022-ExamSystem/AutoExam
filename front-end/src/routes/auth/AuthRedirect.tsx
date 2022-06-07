import React, {SyntheticEvent} from 'react';
import TopNavbar from "../../components/TopNavbar";
import AppLayout from "../../components/AppLayout";
import {Container, Image, Form, Button} from "react-bootstrap";
import {nanoid} from "nanoid";
import autolab_logo from '../../images/autolab_logo.png';

const AutolabLogo = () => {
    //return <img src={autolab_logo}  alt="Autolab Logo"/>
    return <></>
}

const RedirectPage = ({pageLink}: { pageLink: string }) => {
    const handleSubmit = (e: SyntheticEvent) => {
        e.preventDefault();
        const $authCode = document.getElementById("AutolabAuthcode") as HTMLInputElement;
        console.log($authCode.value);
    }

    return (
        <>
            <TopNavbar/>
            <AppLayout>
                <div className="col col-md-8 mx-auto">
                    <h2>Bind with Autolab</h2>
                    <Container className="mt-3 mb-5">
                        <Container className="text-start">
                            <p>Click the banner to visit Autolab authorization page.</p>
                            <p>After the authorization, there will be an authorization code. Paste the code here to
                                finish binding.</p>
                        </Container>
                        <a target="_blank" href={pageLink}>
                            <Image src={autolab_logo} alt="Autolab Logo">
                            </Image>
                        </a>
                    </Container>
                    <Form className="mb-3" onSubmit={e => {
                        handleSubmit(e)
                    }}>
                        <Form.Group controlId="AutolabAuthcode">
                            <Form.Control type="text" placeholder="Enter Auth Code"/>
                        </Form.Group>
                        <Button className="mt-3" variant="primary" type="submit">
                            Submit
                        </Button>
                    </Form>
                </div>
            </AppLayout>
        </>
    )
}

function AuthRedirect() {

    let autolabLink: string = process.env.REACT_APP_AUTOLAB_LOCATION + "/oauth/authorize";

    autolabLink += `?response_type=code&client_id=${process.env.REACT_APP_CLIENT_ID}`;

    const stateValue = nanoid();
    localStorage.setItem('authStateValue', stateValue);
    autolabLink += `&state=` + stateValue;

    if (process.env.REACT_APP_REDIRECT_URI === 'urn:ietf:wg:oauth:2.0:oob') {
        autolabLink += '&redirect_uri=' + encodeURIComponent(process.env.REACT_APP_REDIRECT_URI);
        return (<RedirectPage pageLink={autolabLink}/>);
    }

    const location = window.location.origin;
    let redirect_uri = location + '/oauth-callback';

    redirect_uri = encodeURIComponent(process.env.REACT_APP_REDIRECT_URI || redirect_uri);
    autolabLink += '&redirect_uri=' + redirect_uri;
    window.location.replace(autolabLink);
    return <></>;
}

export default AuthRedirect;