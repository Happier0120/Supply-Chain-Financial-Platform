package com.example.fabric.call;

import java.io.FileReader;
import java.io.IOException;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.security.InvalidKeyException;
import java.security.PrivateKey;
import java.security.cert.CertificateException;
import java.security.cert.X509Certificate;

import com.example.fabric.constant.Submit;
import org.bouncycastle.util.io.pem.PemReader;
import org.hyperledger.fabric.gateway.*;


public class ChaincodeConnect {

    static {
        System.setProperty("org.hyperledger.fabric.sdk.service_discovery.as_localhost", "true");
    }

    // helper function for getting connected to the gateway
    public static Gateway connect() throws Exception {
        // 创建钱包，加载用户密钥
        Path walletPath = Paths.get("/home", "sweeney", "web3j", "fabric-samples", "test-network", "crypto-config", "wallet");
        Wallet wallet = Wallets.newFileSystemWallet(walletPath);
        populateWallet(wallet);
        // 网络连接，加载fabric信息与peer节点密钥
        Path networkConfigPath = Paths.get("/home", "sweeney", "web3j", "fabric-samples", "test-network", "crypto-config", "connection.yaml");
        // 连接建立，后续可用于执行链码
        Gateway.Builder builder = Gateway.createBuilder();
        builder.identity(wallet, "appUser").networkConfig(networkConfigPath).discovery(true);
        return builder.connect();
    }

    public static void populateWallet(Wallet wallet) throws InvalidKeyException, CertificateException, IOException {
        String certPathAdmin = "/home/sweeney/web3j/fabric-samples/test-network/organizations/peerOrganizations/" + Submit.users.get(Submit.currentUser) + ".example.com/ca/ca." + Submit.users.get(Submit.currentUser) + ".example.com-cert.pem";
        String keyPathAdmin = "/home/sweeney/web3j/fabric-samples/test-network/organizations/peerOrganizations/" + Submit.users.get(Submit.currentUser) + ".example.com/ca/priv_sk";
        X509Certificate certAdmin = Identities.readX509Certificate(new PemReader(new FileReader(certPathAdmin)));
        PrivateKey keyAdmin = Identities.readPrivateKey(new PemReader(new FileReader(keyPathAdmin)));
        X509Identity identityAdmin = Identities.newX509Identity(Submit.users.get(Submit.currentUser), certAdmin, keyAdmin);
        wallet.put("admin", identityAdmin);

        String certPath = "/home/sweeney/web3j/fabric-samples/test-network/organizations/peerOrganizations/" + Submit.users.get(Submit.currentUser) + ".example.com/users/User1@" + Submit.users.get(Submit.currentUser) + ".example.com/msp/signcerts/User1@" + Submit.users.get(Submit.currentUser) + ".example.com-cert.pem";
        String keyPath = "/home/sweeney/web3j/fabric-samples/test-network/organizations/peerOrganizations/" + Submit.users.get(Submit.currentUser) + ".example.com/users/User1@" + Submit.users.get(Submit.currentUser) + ".example.com/msp/keystore/priv_sk";
        X509Certificate cert = Identities.readX509Certificate(new PemReader(new FileReader(certPath)));
        PrivateKey key = Identities.readPrivateKey(new PemReader(new FileReader(keyPath)));
        X509Identity identity = Identities.newX509Identity(Submit.users.get(Submit.currentUser), cert, key);
        wallet.put("appUser", identity);
    }

}
