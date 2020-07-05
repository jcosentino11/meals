import React, { useState, useEffect } from 'react';
import { config } from '../config'
import { useAuth0 } from '@auth0/auth0-react';

const Profile = () => {

    const { user, isAuthenticated, getAccessTokenSilently  } = useAuth0();
    const [msg, setMessage] = useState(null);

    useEffect(() => {
        (async () => {
            try {

                if (!isAuthenticated) {
                    console.log("not authenticated yet");
                    return;
                }

                console.log("getting access token...")

                const accessToken = await getAccessTokenSilently({
                    audience: config.auth0.audience,
                    scope: config.auth0.scope,
                });

                const getMessageUrl = `${config.backend.rootUrl}/hello`;

                console.log(`sending request to ${getMessageUrl}`)

                const resp = await fetch(getMessageUrl, {
                    headers: {
                        Authorization: `Bearer ${accessToken}`
                    }
                });

                console.log(`response status: ${resp.status}, content: ${resp.body}`)

                const { msg } = await resp.json();
                
                setMessage(msg);
            } catch (e) {
                console.log(e.message);
            }
        })();
    }, [getAccessTokenSilently]);

    return (
        isAuthenticated ? (
            msg ? (
                <div>
                    <img src={user.picture} alt={user.name} />
                    <h2>{user.name}</h2>
                    <p>{user.email}</p>
                    <h3>Response from API</h3>
                    <pre>{msg}</pre>
                </div>
            ) : (
                <div>
                    "No message found"
                </div>
            )
        ) : (null)
    );
};

export default Profile;