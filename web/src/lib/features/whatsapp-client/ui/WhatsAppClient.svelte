<script lang="ts">
    import { credentials, formFields, responseLog } from "../model/store";
    import { fetchGreenApi } from "../api/client";

    async function callGreenApi(method: string, payload: object | null = null) {
        console.group(`Green API Call: ${method}`);
        console.log("Payload:", payload);

        responseLog.set("Loading...");
        try {
            const result = await fetchGreenApi(method, $credentials, payload);
            responseLog.set(JSON.stringify(result, null, 2));
        } catch (e: unknown) {
            console.error("Error:", e);
            if (e instanceof Error) {
                responseLog.set(`Error: ${e.message}`);
            } else {
                responseLog.set(`An unknown error occurred: ${JSON.stringify(e)}`);
            }
        } finally {
            console.groupEnd();
        }
    }

    $: credentialsReady = !!($credentials.idInstance && $credentials.apiTokenInstance);
</script>

<div class="container">
    <div class="left">
        <section>
            <input bind:value={$credentials.idInstance} placeholder="idInstance" />

            <input
                bind:value={$credentials.apiTokenInstance}
                placeholder="ApiTokenInstance"
            />

            <button
                on:click={() => callGreenApi("getSettings")}
                disabled={!credentialsReady}
                >getSettings</button
            >

            <button
                on:click={() => callGreenApi("getStateInstance")}
                disabled={!credentialsReady}
                >getStateInstance</button
            >
        </section>

        <section>
            <input
                bind:value={$formFields.phoneMessage}
                placeholder="79991234567"
            />

            <input
                bind:value={$formFields.messageText}
                placeholder="Message text"
            />

            <button
                on:click={() =>
                    callGreenApi("sendMessage", {
                        chatId: `${$formFields.phoneMessage}@c.us`,
                        message: $formFields.messageText,
                    })}
                disabled={!credentialsReady || !$formFields.phoneMessage || !$formFields.messageText}
                >sendMessage</button
            >
        </section>

        <section>
            <input
                bind:value={$formFields.phoneFile}
                placeholder="79991234567"
            />
            <input bind:value={$formFields.fileUrl} placeholder="URL файла" />
            <button
                on:click={() =>
                    callGreenApi("sendFileByUrl", {
                        chatId: `${$formFields.phoneFile}@c.us`,
                        urlFile: $formFields.fileUrl,
                        fileName:
                            $formFields.fileUrl.split("/").pop() || "file",
                    })}
                disabled={!credentialsReady || !$formFields.phoneFile || !$formFields.fileUrl}
                >sendFileByUrl</button
            >
        </section>
    </div>

    <div class="right">
        <textarea readonly value={$responseLog}></textarea>
    </div>
</div>
