import { Container, Paper, Typography } from "@material-ui/core";

// Figure we could have some small notes from us on the title pages
export default function DevNotes() {
    return (
        <Container>
            <Paper>
                <Typography>
                    <strong>Dev Notes on V.0.0.Something</strong>
                    <p>
                        Notes:
                        Now does shit
                    </p>
                </Typography>
            </Paper>
        </Container>
     );
}