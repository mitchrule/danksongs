import styles from "../styles/Navbar.module.css";

export default function Navbar() {
  return (
    <>
      <div className={styles.navbar}>
        <h1>DankSongs</h1>
        <ul>
          <li>
            <a href="/">Home</a>
          </li>
          <li>
            <a href="/playlists">Playlists</a>
          </li>
          <li>
            <a href="/account">Account</a>
          </li>
          <li style={{ float: "right" }}>
            <a href="/logout">Sign Out</a>
          </li>
        </ul>
      </div>
    </>
  );
}
