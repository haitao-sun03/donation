import React from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  faChartLine,
  faShieldAlt,
  faHandHoldingUsd,
} from "@fortawesome/free-solid-svg-icons";
import "@fortawesome/fontawesome-svg-core/styles.css";
import { config } from "@fortawesome/fontawesome-svg-core";
config.autoAddCss = false;
import "./DonationHome.css";

const DonationHome = () => {
  return (
    <div className="donation-home">
      <div className="hero-section">
        <h1>Welcome to AquaSeed , A Web3 Donation Platform</h1>
        <p>
          Connect your wallet to fund campaigns with AquaCoin and earn EcoSeed
          NFTs.
        </p>
      </div>

      <div className="features-section">
        <div className="feature-card">
          <div className="icon-container">
            <FontAwesomeIcon icon={faChartLine} className="feature-icon" />
          </div>
          <h3>Transparent Donations</h3>
          <p>All donation records are stored on the blockchain</p>
        </div>
        <div className="feature-card">
          <div className="icon-container">
            <FontAwesomeIcon icon={faShieldAlt} className="feature-icon" />
          </div>
          <h3>Secure & Reliable</h3>
          <p>Using smart contracts to ensure fund safety</p>
        </div>
        <div className="feature-card">
          <div className="icon-container">
            <FontAwesomeIcon icon={faHandHoldingUsd} className="feature-icon" />
          </div>
          <h3>Easy to Use</h3>
          <p>Simple steps to complete donations</p>
        </div>
      </div>
    </div>
  );
};

export default DonationHome;
