Name:       telegramnotify
Version:    %{version}
Release:    %{release}
Summary:    User friendly cross platform console application to send notifications (text, images, files) via telegram bots in proper way.
License:    MIT
URL:        https://github.com/vodolaz095/telegramnotify/


%description
User friendly cross platform console application to send notifications (text, images, files) via telegram bots in proper way.

%prep

%build

%check

%install
mkdir -p %{buildroot}/usr/bin/
mkdir -p %{buildroot}/etc/
install -m 755 telegramnotify %{buildroot}/usr/bin/telegramnotify
install -m 644 telegramnotify.json %{buildroot}/etc/telegramnotify.json

%files
%defattr(-,root,root)
/usr/bin/telegramnotify
/etc/telegramnotify.json
