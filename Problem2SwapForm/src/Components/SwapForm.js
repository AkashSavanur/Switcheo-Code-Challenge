import { useState, useEffect } from "react";
import axios from "axios";
import Button from "@mui/material/Button";
import Swal from "sweetalert2";
import SwapHorizontalCircleIcon from "@mui/icons-material/SwapHorizontalCircle";
import IconButton from "@mui/material/IconButton";
import LightModeIcon from "@mui/icons-material/LightMode";
import DarkModeIcon from "@mui/icons-material/DarkMode";

export default function SwapForm() {
  const [tokens, setTokens] = useState([]);
  const [fromToken, setFromToken] = useState("USD");
  const [toToken, setToToken] = useState("ETH");
  const [amount, setAmount] = useState("");
  const [exchangeRate, setExchangeRate] = useState(null);
  const [loading, setLoading] = useState(false);
  const [message, setMessage] = useState("");

  useEffect(() => {
    axios
      .get("https://interview.switcheo.com/prices.json")
      .then((response) => {
        const uniqueTokens = [
          ...new Set(response.data.map((item) => item.currency)),
        ];
        setTokens(uniqueTokens);
      })
      .catch(() => setTokens([]));
  }, []);

  useEffect(() => {
    if (fromToken && toToken) {
      axios
        .get("https://interview.switcheo.com/prices.json")
        .then((response) => {
          const fromPrice = response.data.find(
            (item) => item.currency === fromToken
          )?.price;
          const toPrice = response.data.find(
            (item) => item.currency === toToken
          )?.price;
          if (fromPrice && toPrice) {
            setExchangeRate(toPrice / fromPrice);
          } else {
            setExchangeRate(null);
          }
        })
        .catch(() => setExchangeRate(null));
    }
  }, [fromToken, toToken]);

  const handleSwap = () => {
    if (!amount || isNaN(amount) || amount <= 0) {
      Swal.fire({
        title: "Invalid!",
        text: "Please enter a valid amount",
        icon: "error",
      });
      return;
    } else if (!fromToken || !toToken) {
      Swal.fire({
        title: "Invalid!",
        text: "Please enter a currency",
        icon: "error",
      });
      return;
    } else {
      setLoading(true);
      setTimeout(() => {
        setLoading(false);
        setMessage(
          `${amount} ${fromToken} = ${(amount * exchangeRate).toFixed(
            2
          )} ${toToken}`
        );
      }, 500);
    }
  };

  const handleReset = () => {
    setAmount("");
    setFromToken("");
    setToToken("");
    setMessage("");
    setExchangeRate(null);
  };

  const handleReverse = () => {
    setToToken(fromToken);
    setFromToken(toToken);
  };

  return (
    <div
      style={{ display: "flex", flexDirection: "column", minHeight: "100vh" }}
    >
      <header
        style={{
          backgroundColor: "#00002a",
          color: "white",
          textAlign: "center",
          padding: "16px",
          fontSize: "24px",
          fontWeight: "bold",
        }}
      >
        <span>Currency Swap Platform</span>
      </header>

      <main
        style={{
          flex: 1,
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          backgroundColor: "#000053",
          padding: "16px",
        }}
      >
        <div
          style={{
            maxWidth: "800px",
            width: "100%",
            backgroundColor: "white",
            padding: "16px",
            borderRadius: "16px",
            boxShadow: "0px 4px 6px rgba(0, 0, 0, 0.1)",
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
          }}
        >
          <h2
            style={{
              fontSize: "30px",
              fontWeight: "bold",
              marginBottom: "16px",
            }}
          >
            Swap Your Currency
          </h2>
          <div
            style={{
              display: "flex",
              gap: "8px",
              width: "100%",
              marginBottom: "8px",
              alignItems: "flex-start",
            }}
          >
            <div style={{ width: "100%" }}>
              <select
                style={{
                  width: "100%",
                  padding: "8px",
                  border: "1px solid #ccc",
                  borderRadius: "8px",
                  textAlign: "center",
                }}
                value={fromToken}
                onChange={(e) => setFromToken(e.target.value)}
              >
                <option value="">From</option>
                {tokens.map((token) => (
                  <option key={token} value={token}>
                    {token}
                  </option>
                ))}
              </select>
              {fromToken && (
                <img
                  src={require(`../Logos/${fromToken.toUpperCase()}.svg`)}
                  alt={fromToken}
                  style={{
                    width: "50px",
                    height: "50px",
                    marginTop: "8px",
                    display: "block",
                    marginLeft: "auto",
                    marginRight: "auto",
                  }}
                />
              )}
            </div>
            {fromToken && toToken && (
              <div style={{ display: "flex", alignItems: "center" }}>
                <IconButton onClick={handleReverse}>
                  <SwapHorizontalCircleIcon
                    fontSize="large"
                    color="secondary"
                  />
                </IconButton>
              </div>
            )}
            <div style={{ width: "100%" }}>
              <select
                style={{
                  width: "100%",
                  padding: "8px",
                  border: "1px solid #ccc",
                  borderRadius: "8px",
                  textAlign: "center",
                }}
                value={toToken}
                onChange={(e) => setToToken(e.target.value)}
              >
                <option value="">To</option>
                {tokens.map((token) => (
                  <option key={token} value={token}>
                    {token}
                  </option>
                ))}
              </select>
              {toToken && (
                <img
                  src={require(`../Logos/${toToken.toUpperCase()}.svg`)}
                  alt={toToken}
                  style={{
                    width: "50px",
                    height: "50px",
                    marginTop: "8px",
                    display: "block",
                    marginLeft: "auto",
                    marginRight: "auto",
                  }}
                />
              )}
            </div>
          </div>

          <input
            type="number"
            style={{
              width: "30%",
              padding: "8px",
              border: "1px solid #ccc",
              borderRadius: "8px",
              textAlign: "center",
              marginTop: "20px",
              marginBottom: "20px",
            }}
            placeholder="Enter An Amount"
            value={amount}
            onChange={(e) => setAmount(e.target.value)}
          />
          <div
            style={{
              color: "#000",
              textAlign: "center",
              marginBottom: "8px",
            }}
          >
            Exchange Rate: {exchangeRate ? exchangeRate.toFixed(4) : ""}
          </div>
          <div
            style={{
              display: "flex",
              gap: "8px",
              width: "100%",
              marginBottom: "8px",
            }}
          >
            <Button
              style={{
                width: "100%",
                backgroundColor: "grey",
                color: "white",
                padding: "12px",
                borderRadius: "8px",
              }}
              onClick={handleReset}
            >
              Reset
            </Button>
            <Button
              style={{
                width: "100%",
                backgroundColor: "#00002a",
                color: "white",
                padding: "12px",
                borderRadius: "8px",
              }}
              onClick={handleSwap}
              disabled={loading}
            >
              {loading ? "Converting..." : "Convert"}
            </Button>
          </div>
          {message && (
            <div
              style={{
                marginTop: "16px",
                textAlign: "center",
                color: "green",
                fontSize: "20px",
              }}
            >
              {message}
            </div>
          )}
        </div>
      </main>
      <footer
        style={{
          backgroundColor: "#00002a",
          color: "white",
          textAlign: "center",
          padding: "16px",
          fontSize: "14px",
        }}
      >
        &copy; 2025 Swap Platform. Created by Akash Savanur.
      </footer>
    </div>
  );
}
