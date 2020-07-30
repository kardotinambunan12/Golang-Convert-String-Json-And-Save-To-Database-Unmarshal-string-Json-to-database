-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 30 Jul 2020 pada 04.47
-- Versi server: 10.4.11-MariaDB
-- Versi PHP: 7.3.16

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `user`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `brands`
--

CREATE TABLE `brands` (
  `id` int(11) NOT NULL,
  `business_id` int(11) DEFAULT NULL,
  `mobile_id` varchar(50) DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL,
  `description` text DEFAULT NULL,
  `created_by` int(11) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `brands`
--

INSERT INTO `brands` (`id`, `business_id`, `mobile_id`, `name`, `description`, `created_by`, `created_at`, `updated_at`) VALUES
(6, 1, '2222', 'Ajinamoto ', 'Bumbu Masakan Dapur', 1, '2020-07-29 09:09:45', NULL),
(7, 1, '1111', 'Masako', 'Bumbu Masakan Dapur', 1, '2020-07-29 09:10:21', NULL);

-- --------------------------------------------------------

--
-- Struktur dari tabel `tb_product`
--

CREATE TABLE `tb_product` (
  `id` int(12) UNSIGNED NOT NULL,
  `name` varchar(100) CHARACTER SET armscii8 NOT NULL,
  `type` varchar(50) CHARACTER SET armscii8 NOT NULL,
  `category` varchar(120) CHARACTER SET armscii8 NOT NULL,
  `image` varchar(120) CHARACTER SET armscii8 NOT NULL,
  `description` varchar(200) CHARACTER SET armscii8 DEFAULT NULL,
  `harga` decimal(20,0) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Struktur dari tabel `tb_user`
--

CREATE TABLE `tb_user` (
  `id` int(11) NOT NULL,
  `first_name` varchar(150) CHARACTER SET armscii8 NOT NULL,
  `last_name` varchar(150) CHARACTER SET armscii8 NOT NULL,
  `alamat` varchar(150) CHARACTER SET armscii8 NOT NULL,
  `zip_code` varchar(10) CHARACTER SET koi8r NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `tb_user`
--

INSERT INTO `tb_user` (`id`, `first_name`, `last_name`, `alamat`, `zip_code`) VALUES
(1, 'Timbul', 'Munthe', 'RantauPrapat', '22447'),
(2, 'Hafsah', 'Anugrah', 'Deli Serdang / Medan', '23221'),
(3, 'Kardo', 'Parna', 'Medan / Tapanuli', '22456'),
(6, 'batak', 'pok', 'samosir', ''),
(7, 'kardo', 'pop', 'padang', '1111111'),
(8, 'batak', 'pok', 'samosir', '22456'),
(9, 'batak', 'pok', 'samosir', '22456'),
(10, 'batak', 'pok', 'samosir', '22456');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `brands`
--
ALTER TABLE `brands`
  ADD PRIMARY KEY (`id`) USING BTREE,
  ADD KEY `business_id` (`business_id`),
  ADD KEY `create_by` (`created_by`) USING BTREE;

--
-- Indeks untuk tabel `tb_product`
--
ALTER TABLE `tb_product`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `tb_user`
--
ALTER TABLE `tb_user`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `brands`
--
ALTER TABLE `brands`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT untuk tabel `tb_user`
--
ALTER TABLE `tb_user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
